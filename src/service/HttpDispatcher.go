package service

import (
	"github.com/valyala/fasthttp"
	"errors"
	"strconv"
	"generic-op/utils"
)

type RequestHandler func(ent *EntRequest) (code int, err error, desc string, data []byte)

var reflectionTable map[string]RequestHandler

func init() {
	reflectionTable = make(map[string]RequestHandler)
	reflectionTable["register"] = reqChannel_register
	reflectionTable["login"] = reqChannel_login
}

/*
Get the code and result, response for client
 */
func HttpHandle(req *fasthttp.RequestCtx) () {
	code, err, desc, data := httpService(req)
	req.Response.Header.Set("code", strconv.Itoa(code))
	if err != nil {
		req.Response.Header.Set("err", err.Error())
	} else {
		req.Response.Header.Set("err", "")
	}
	req.Response.Header.Set("desc", desc)
	if data != nil {
		req.Response.SetBody(data)
	}
}

/*
Process the request
 */
func httpService(req *fasthttp.RequestCtx) (code int, err error, desc string, data []byte) {

	// Create ent
	ent := NewEntRequest(req)

	// Check
	err, desc = ent.Check()
	if err != nil {
		return ENT_CHECK_ERROR, err, desc, nil
	}

	// Dispatch due to request type
	h, ok := reflectionTable[ent.Header.ReqType]
	if !ok {
		return INVALID_CHANNEL, errors.New(INVALID_CHANNEL_MSG), INVALID_CHANNEL_MSG, nil
	}
	code, err, desc, data = h(ent)
	return code, err, desc, data
}

/*
Channel to register user data
 */
func reqChannel_register(ent *EntRequest) (code int, err error, desc string, data []byte) {
	if ent.Header.ContentType != "pwd64" {
		return ILLEGAL_CNT_TYPE, errors.New(ILLEGAL_CNT_TYPE_MSG), ILLEGAL_CNT_TYPE_MSG, nil
	}
	pwdBytes, err := utils.Base64DecodeStringToBytes(string(ent.Data.Content))
	if err != nil {
		return ILLEGAL_BASE64_DATA, err, err.Error(), nil
	}
	err = RequestSqlForRegister(ent.Header.UserName, string(pwdBytes))
	if err != nil {
		return REGISTER_ERROR, err, err.Error(), nil
	}
	// todo add user data in db
	return 0, nil, "", nil
}

/*
Channel to login
 */
func reqChannel_login(ent *EntRequest) (code int, err error, desc string, data []byte) {
	if ent.Header.ContentType != "chksmd5" {
		return ILLEGAL_CNT_TYPE, errors.New(ILLEGAL_CNT_TYPE_MSG), ILLEGAL_CNT_TYPE_MSG, nil
	}
	userData, err := RequestSqlForUserData(ent.Header.UserName)
	if err != nil {
		return GET_USER_DATA_ERROR, err, err.Error(), nil
	}
	checksumMd5 := string(ent.Data.Content)
	checksumExpected := userData.Password + ent.Header.TimeStamp + userData.UserName
	checksumExpectedMd5 := utils.Md5Encode(checksumExpected)
	if checksumMd5 == checksumExpectedMd5 {
		return SUCCESS, nil, "", nil
	}
	return INCORRECT_PASSWORD, errors.New(INCORRECT_PASSWORD_MSG), INCORRECT_PASSWORD_MSG, nil
}
