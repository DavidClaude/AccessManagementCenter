package service

import (
	"github.com/valyala/fasthttp"
	"strings"
	"errors"
)

type EntRequest struct {
	Header *EntRequestHeader `json:"header"`
	Data   *EntRequestData   `json:"data"`
}

type EntRequestHeader struct {
	ReqType string `json:"type"`
	TimeStamp string `json:"timestamp"`
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

func (er *EntRequest)Check()(err error, desc string)  {

	// Check request type
	req_type := er.Header.ReqType
	b, err := Config.GetBool("request_type", req_type)
	if err != nil {
		return  err, "get request type error"
	}
	if !b {
		return errors.New("unavailable type"), "request type is unavailable"
	}
	return nil, ""

	// Check data
	l := len(er.Data.Content)
	if (l <= 0){
		return errors.New("empty data"), "data is empty"
	}
}
