package hc

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleHc(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/hc", nil)

	handle().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	resp := w.Result()
	require.Equal(t, 200, resp.StatusCode, "want response code 200")
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}
