package game

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"tracing/pkg/encrypt"
)

type serv struct {
	client  *resty.Client
	host    string
	source  int
	signKey string
}

func (serv *serv) request(ctx context.Context, query Encryptor, data any, path string) {

	var resp Response
	resp.Data = data
	query.source(serv.source)
	query.signKey(serv.signKey)

	var resultErr error
	respRaw, err := serv.client.R().
		SetQueryString(query.String() + "&sign=" + encrypt.MD5(query.Encode())).
		SetResult(&resp).
		SetError(&resultErr).
		Get(serv.host + path)
	fmt.Println(respRaw)
	fmt.Println(err)
	fmt.Println(resultErr)

}
