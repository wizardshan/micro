package game

import (
	"context"
	"tracing/domain/money"
	"tracing/pkg/app"
	"tracing/pkg/http"
)

type Payment struct {
	serv
}

func NewPayment(client *http.Request, components *app.Components) *Payment {
	conf := components.Servers.Payment
	return &Payment{
		serv{client, conf.Host, conf.Source, conf.SignKey},
	}
}

func (serv *Payment) Orders(ctx context.Context, queryFunc func(q *OrderQuery)) UserOnlineTimes {

	query := &OrderQuery{
		MoneyType: OrderMoneyTypeDurationList,
		MinMoney:  money.Fen * 0,
	}
	queryFunc(query)

	var data UserOnlineTimes
	serv.request(ctx, query, &data, "/order/PropOrder/getUserOrderMoney")
	return data
}
