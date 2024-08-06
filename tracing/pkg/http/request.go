package http

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Request struct {
	client *resty.Client
}

func New(client *resty.Client) *Request {
	req := new(Request)
	req.client = resty.New()
	return req
}

func (request *Request) Get(ctx context.Context, values any, url string) {
	q, err := query.Values(values)
	if err != nil {
		panic(err)
	}
	fmt.Println(q.Encode())

	resp, err := request.client.R().
		SetQueryString(q.Encode()).
		Get(url)
	fmt.Println(resp, err)
}
