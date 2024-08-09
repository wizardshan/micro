package user

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	host   string
	salt   string
	client *resty.Client
}

func New(host string, client *resty.Client) *Client {
	c := new(Client)
	c.host = host
	c.salt = "youxuIlop6JD8ftf7BdPLxi"
	c.client = client

	return c
}

//func (client *Client) Fetch(ctx context.Context, id int) *domain.User {
//	return repo.db.User.Query().WithComments().WithPosts().Where(user.ID(id)).FirstX(ctx).Mapper()
//}

func (c *Client) FetchMany(ctx context.Context, IDs []int) Users {
	//query := &QueryUsers{
	//	Query:   defaultQuery(),
	//	UserIDs: pie.Join(IDs, ","),
	//}

	type Resp struct {
		Response
		Data UserByIDMapping `json:"data"`
	}

	resp := new(Resp)
	//c.request(ctx, query, resp, "/apis/inner/info/batch/byIds.action")
	fmt.Println(resp)
	return nil
}

//func (c *Client) request(ctx context.Context, value any, resp any, path string) Users {
//	query := ""
//	if builder, ok := value.(http.Builder); ok {
//		query = builder.Build() + "&sign=" + encrypt.Sign(value, c.salt)
//	}
//
//	c.requestClient.Get(ctx, query, resp, c.host+path)
//	return nil
//}
