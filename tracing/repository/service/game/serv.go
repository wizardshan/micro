package game

import (
	"context"
	"tracing/pkg/encrypt"
	"tracing/pkg/http"
)

type serv struct {
	client  *http.Request
	host    string
	source  int
	signKey string
}

func (serv *serv) request(ctx context.Context, query http.Query, data any, path string) {

	var resp Response
	resp.Data = data

	s, ok := query.(encrypt.Signer)
	if !ok {
		panic("value must be signer")
	}
	s.Build(serv.source, serv.signKey)
	serv.client.Get(ctx, query.String()+"&sign="+encrypt.MD5(s.Encode()), &resp, serv.host+path)
}
