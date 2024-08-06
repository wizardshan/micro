package usercenter

import (
	"context"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"tracing/pkg/http"
)

type Users []*User

type User struct {
}

type Client struct {
	host    string
	request *http.Request
}

func New(host string, client *http.Request) *Client {
	c := new(Client)
	c.host = host
	c.request = client

	return c
}

//func (client *Client) Fetch(ctx context.Context, id int) *domain.User {
//	return repo.db.User.Query().WithComments().WithPosts().Where(user.ID(id)).FirstX(ctx).Mapper()
//}

func (c *Client) FetchMany(ctx context.Context, IDs []int) Users {
	query := &QueryUser{
		Query:   defaultQuery(),
		UserIDs: pie.Join(IDs, ","),
	}

	c.request.Get(ctx, query, fmt.Sprintf("%s%s", c.host, "/apis/inner/info/batch/byIds.action"))
	return nil
}
