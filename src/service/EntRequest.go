package service

import (
	"github.com/valyala/fasthttp"
	"strings"
	"errors"
	"strconv"
	"time"
	"generic-op/utils"
)

type EntRequest struct {
	Header *EntRequestHeader `json:"header"`
	Data   *EntRequestData   `json:"data"`
}

type EntRequestHeader struct {
	ReqType   string `json:"req_type"`
	TimeStamp string `json:"time_stamp"`
	UserName  string `json:"user_name"`
}

type EntRequestData struct {
	Content []byte `json:"content"`
}

func NewEntRequest(req *fasthttp.RequestCtx) (ent *EntRequest) {
	header := &EntRequestHeader{
		ReqType: strings.TrimSpace(string(req.Request.Header.Peek("type"))),
	}
	body := req.PostBody()
	data := &EntRequestData{
		Content: body,
	}
	ent = &EntRequest{
		Header: header,
		Data:   data,
	}
	return ent
}

func (er *EntRequest) Check() (err error, desc string) {

	// Check request type
	req_type := er.Header.ReqType
	b, err := Config.GetBool("request_type", req_type)
	if err != nil {
		return err, err.Error()
	}
	if !b {
		return errors.New(ENT_CHECK_ERROR_MSG), "req_type is unavailable"
	}

	// Check timestamp
	period, err := Config.GetInt64("check", "valid_period")
	if err != nil {
		return err, err.Error()
	}
	tsInt64, err := strconv.ParseInt(er.Header.TimeStamp, 10, 64)
	if err != nil {
		return err, err.Error()
	}
	tsOS := time.Now().Unix()
	intv := utils.AbsInt64(tsOS - tsInt64)
	if intv > period {
		return errors.New(ENT_CHECK_ERROR_MSG), "request is expired"
	}

	// Check data
	l := len(er.Data.Content)
	if (l <= 0) {
		return errors.New(ENT_CHECK_ERROR_MSG), "data is empty"
	}
	return nil, ""
}
