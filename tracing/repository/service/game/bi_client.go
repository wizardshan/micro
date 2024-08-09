package game

import (
	"context"
	"tracing/pkg/app"
	"tracing/pkg/http"
)

type BI struct {
	serv
}

func NewBI(client *http.Request, components *app.Components) *BI {
	conf := components.Servers.BI
	return &BI{
		serv{client, conf.Host, conf.Source, conf.SignKey},
	}
}

func (serv *BI) OnlineTimes(ctx context.Context, queryFunc func(q *OnlineTimeQuery)) UserOnlineTimes {
	query := serv.defaultOnlineTimeQuery()

	var data UserOnlineTimes
	serv.request(ctx, query, &data, "/mixed/MOnlineTime/lists")
	return data
}

func (serv *BI) terminal() []int {
	return []int{1, 3, 4, 5, 6, 7}
}

func (serv *BI) defaultOnlineTimeQuery() *OnlineTimeQuery {
	return &OnlineTimeQuery{
		Terminal: serv.terminal(),
	}
}
