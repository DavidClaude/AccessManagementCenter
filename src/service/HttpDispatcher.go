package service

import (
	"github.com/valyala/fasthttp"
	"errors"
)

func HttpHandle(req *fasthttp.RequestCtx) () {
	code, err, desc, result := httpService(req)
	if (code == 0) {
		// todo return success
	} else {
		// todo return error
	}
}

func httpService(req *fasthttp.RequestCtx) (code int, err error, desc string, result string) {

	// New param
	ent := NewEntRequest(req)

	// Check common info
	err, desc = ent.Check()
	if err != nil {
		return CHECK_ENT_REQ_ERROR, err, desc, ""
	}

	// Dispatch due to request type
	switch ent.Header.ReqType {
	case "apply":
		return reqChannel_apply(ent)
	case "login":
		return reqChannel_login(ent)
	default:
		return INVALID_CHANNEL, errors.New("invalid channel"), "channel is valid", ""
	}

}

func reqChannel_apply(ent *EntRequest) (code int, err error, desc string, result string) {

}

func reqChannel_login(ent *EntRequest) (code int, err error, desc string, result string) {

}
