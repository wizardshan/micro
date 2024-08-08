package bi

import (
	"context"
	"tracing/pkg/app"
	"tracing/pkg/encrypt"
	"tracing/pkg/http"
)

type Client struct {
	requestClient *http.Request
	host          string
	source        int
	signKey       string
}

func New(requestClient *http.Request, components *app.Components) *Client {
	c := new(Client)
	c.host = components.Servers.BI.Host
	c.source = components.Servers.BI.Source
	c.signKey = components.Servers.BI.SignKey
	c.requestClient = requestClient

	return c
}

func (c *Client) OnlineTimes(ctx context.Context, query *OnlineTimeQuery) UserOnlineTimes {
	var data UserOnlineTimes
	c.request(ctx, query, &data, "/mixed/MOnlineTime/lists")
	return data
}

func (c *Client) request(ctx context.Context, query http.Query, data any, path string) {

	var resp Response
	resp.Data = data

	s, ok := query.(encrypt.Signer)
	if !ok {
		panic("value must be signer")
	}
	s.Build(c.source, c.signKey)

	c.requestClient.Get(ctx, query.String()+"&sign="+encrypt.MD5(s.Encode()), &resp, c.host+path)
}
