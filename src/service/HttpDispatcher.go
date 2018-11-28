package service

import (
	"github.com/valyala/fasthttp"
	"errors"
	"strconv"
	"generic-op/utils"
)

type RequestHandler func(ent *EntRequest) (code int, err error, desc string)

var reflectionTable map[string]RequestHandler

func init() {
	reflectionTable = make(map[string]RequestHandler)
	reflectionTable["apply"] = reqChannel_apply
	reflectionTable["login"] = reqChannel_login
}

/*
Get the code and result, response for client
 */
func HttpHandle(req *fasthttp.RequestCtx) () {
	code, err, desc := httpService(req)
	req.Response.Header.Set("code", strconv.Itoa(code))
	if err != nil {
		req.Response.Header.Set("err", err.Error())
	} else {
		req.Response.Header.Set("err", "")
	}

	req.Response.Header.Set("desc", desc)
}

/*
Process the request
 */
func httpService(req *fasthttp.RequestCtx) (result_code int, err error, desc string) {

	// Create ent
	ent := NewEntRequest(req)

	// Check
	err, desc = ent.Check()
	if err != nil {
		return ENT_CHECK_ERROR, err, desc
	}

	// Dispatch due to request type
	h, ok := reflectionTable[ent.Header.ReqType]
	if !ok {
		return INVALID_CHANNEL, errors.New(INVALID_CHANNEL_MSG), INVALID_CHANNEL_MSG
	}
	result_code, err, desc = h(ent)
	return result_code, err, desc
}

/*
Channel to apply for user data
 */
func reqChannel_apply(ent *EntRequest) (code int, err error, desc string) {
	return 0, nil, ""
}

/*
Channel to login
 */
func reqChannel_login(ent *EntRequest) (code int, err error, desc string) {
	userData, err := GetUserDataFromSql(ent.Header.UserName)
	if err != nil {
		return GET_USER_DATA_ERROR, err, err.Error()
	}
	checksumMd5 := string(ent.Data.Content)
	checksumExpected := userData.Password + ent.Header.TimeStamp + userData.UserName
	checksumExpectedMd5 := utils.Md5Encode(checksumExpected)
	if checksumMd5 == checksumExpectedMd5 {
		return SUCCESS, nil, ""
	}
	return INCORRECT_PASSWORD, errors.New(INCORRECT_PASSWORD_MSG), INCORRECT_PASSWORD_MSG
}
