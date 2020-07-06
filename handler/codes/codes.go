package codes

import (
	"strconv"

	"github.com/twitchtv/twirp"
)

const (
	Code = "code"

	// Custom Err code
	InvalidArguments = 1000

	BookNameTooShort = 1001
)

func With(err error, code int) error {
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	return twerr.WithMeta(Code, strconv.Itoa(code))
}

func Get(err error) int {
	terr, ok := err.(twirp.Error)
	if !ok {
		terr = twirp.InternalErrorWith(err)
	}

	if code, err := strconv.Atoi(terr.Meta(Code)); err == nil {
		return code
	}

	terrCode := terr.Code()
	switch terrCode {
	case twirp.InvalidArgument:
		return InvalidArguments
	default:
		return twirp.ServerHTTPStatusFromErrorCode(terrCode)
	}
}
