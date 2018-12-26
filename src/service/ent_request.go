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
	ReqType     string `json:"req_type"`
	TimeStamp   string `json:"time_stamp"`
	UserName    string `json:"user_name"`
	ContentType string `json:"content_type"`
}

type EntRequestData struct {
	Content []byte `json:"cnt"`
}

func NewEntRequest(req *fasthttp.RequestCtx) (ent *EntRequest) {
	header := &EntRequestHeader{
		ReqType:     strings.TrimSpace(string(req.Request.Header.Peek("req_type"))),
		TimeStamp:   strings.TrimSpace(string(req.Request.Header.Peek("time_stamp"))),
		UserName:    strings.TrimSpace(string(req.Request.Header.Peek("user_name"))),
		ContentType: strings.TrimSpace(string(req.Request.Header.Peek("content_type"))),
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
	b, err := ConfigIns.GetBool("request_type", req_type)
	if err != nil {
		return err, err.Error()
	}
	if !b {
		return errors.New(ENT_CHECK_ERROR_MSG), "request type is unavailable"
	}

	// Check timestamp
	period, err := ConfigIns.GetInt64("local", "valid_period")
	if err != nil {
		return err, err.Error()
	}
	tsInt64, err := strconv.ParseInt(er.Header.TimeStamp, 10, 64)
	if err != nil {
		return err, err.Error()
	}
	tsOSInt64 := time.Now().Unix()
	intv := utils.AbsInt64(tsOSInt64 - tsInt64)
	if intv > period {
		return errors.New(ENT_CHECK_ERROR_MSG), "request is expired"
	}

	// Check content type
	cnt_type := er.Header.ContentType
	b, err = ConfigIns.GetBool("content_type", cnt_type)
	if err != nil {
		return err, err.Error()
	}
	if !b {
		return errors.New(ENT_CHECK_ERROR_MSG), "content type is unavailable"
	}

	// Check data
	l := len(er.Data.Content)
	if (l <= 0) {
		return errors.New(ENT_CHECK_ERROR_MSG), "data is empty"
	}
	return nil, ""
}
