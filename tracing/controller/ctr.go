package controller

import "encoding/json"

type Response struct {
	Code    int
	Message string
	Success bool
	Router  string
	Cache   bool
	Data    any
}

func (resp *Response) bytes() []byte {
	b, _ := json.Marshal(resp)
	return b
}
