package render

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/oxtoacart/bpool"
	"github.com/twitchtv/twirp"
	"github.com/yiplee/twirp-gateway-example/handler/codes"
)

type wrapResponse struct {
	wrapData bool
	writer   http.ResponseWriter
	header   http.Header
	status   int
	pool     *bpool.BufferPool
}

func (w *wrapResponse) Header() http.Header {
	return w.header
}

func (w *wrapResponse) WriteHeader(statusCode int) {
	w.status = statusCode
}

func (w *wrapResponse) finalWrite(body []byte) (int, error) {
	// reset content length
	w.header.Set("Content-Length", strconv.Itoa(len(body)))

	for key := range w.header {
		w.writer.Header().Set(key, w.header.Get(key))
	}

	w.writer.WriteHeader(w.status)
	return w.writer.Write(body)
}

func (w *wrapResponse) isJsonContent() bool {
	typ := w.header.Get("Content-Type")
	return strings.HasPrefix(typ, "application/json")
}

func (w *wrapResponse) Write(body []byte) (int, error) {
	switch w.status {
	case http.StatusOK:
		if w.wrapData && w.isJsonContent() {
			return w.writeData(body)
		} else {
			return w.writeRaw(body)
		}
	default:
		return w.writeErr(body)
	}
}

func (w *wrapResponse) writeRaw(body []byte) (int, error) {
	return w.finalWrite(body)
}

type dataResponse struct {
	Data json.RawMessage `json:"data,omitempty"`
}

func (w *wrapResponse) writeData(body []byte) (int, error) {
	b := w.pool.Get()
	defer w.pool.Put(b)

	r := dataResponse{Data: body}
	if err := json.NewEncoder(b).Encode(r); err != nil {
		return 0, err
	}

	return w.finalWrite(b.Bytes())
}

type errorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (w *wrapResponse) writeErr(body []byte) (int, error) {
	r := errorResponse{
		Code: w.status,
		Msg:  http.StatusText(w.status),
	}

	var tj twerrJSON
	if err := json.Unmarshal(body, &tj); err == nil {
		twerr := twirp.NewError(twirp.ErrorCode(tj.Code), tj.Msg)
		if tj.Meta != nil {
			if code, _ := tj.Meta[codes.Code]; code != "" {
				twerr = twerr.WithMeta(codes.Code, code)
			}
		}

		r.Code = codes.Get(twerr)
		r.Msg = twerr.Msg()
	}

	b, err := json.Marshal(r)
	if err != nil {
		return 0, err
	}

	return w.finalWrite(b)
}

func WrapResponse(wrapData bool) func(http.Handler) http.Handler {
	bufferPool := bpool.NewBufferPool(64)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w = &wrapResponse{
				writer:   w,
				header:   http.Header{},
				wrapData: wrapData,
				pool:     bufferPool,
			}

			next.ServeHTTP(w, r)
		})
	}
}
