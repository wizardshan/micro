package main

import (
	"github.com/go-resty/resty/v2"
)

func initClientHttp() *resty.Client {
	return resty.New()
}
