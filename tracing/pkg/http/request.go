package http

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Query interface {
	String() string
}

type Request struct {
	client *resty.Client
}

func New(client *resty.Client) *Request {
	req := new(Request)
	req.client = resty.New()
	return req
}

func (request *Request) Get(ctx context.Context, query string, result any, url string) {

	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	var resultErr error
	resp, err := request.client.R().
		SetQueryString(query).
		SetResult(result).
		SetError(&resultErr).
		Get(url)
	fmt.Println(resp)
	fmt.Println(err)
	fmt.Println(resultErr)

}
