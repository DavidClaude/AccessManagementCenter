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
	ReqType     string `json:"req_tp"`
	TimeStamp   string `json:"ts"`
	UserName    string `json:"usr"`
	ContentType string `json:"cnt_tp"`
}

type EntRequestData struct {
	Content []byte `json:"cnt"`
}

func NewEntRequest(req *fasthttp.RequestCtx) (ent *EntRequest) {
	header := &EntRequestHeader{
		ReqType:     strings.TrimSpace(string(req.Request.Header.Peek("req_tp"))),
		TimeStamp:   strings.TrimSpace(string(req.Request.Header.Peek("ts"))),
		UserName:    strings.TrimSpace(string(req.Request.Header.Peek("usr"))),
		ContentType: strings.TrimSpace(string(req.Request.Header.Peek("cnt_tp"))),
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
	b, err := Config.GetBool("req_tp", req_type)
	if err != nil {
		return err, err.Error()
	}
	if !b {
		return errors.New(ENT_CHECK_ERROR_MSG), "req_tp is unavailable"
	}

	// Check timestamp
	period, err := Config.GetInt64("local", "valid_period")
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
	b, err = Config.GetBool("cnt_tp", cnt_type)
	if err != nil {
		return  err, err.Error()
	}
	if !b {
		return errors.New(ENT_CHECK_ERROR_MSG), "cnt_tp is unavailable"
	}

	// Check data
	l := len(er.Data.Content)
	if (l <= 0) {
		return errors.New(ENT_CHECK_ERROR_MSG), "data is empty"
	}
	return nil, ""
}
