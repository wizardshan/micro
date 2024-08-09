package game

import (
	"fmt"
	"testing"
	"tracing/pkg/query"
)

func TestOnlineTimeQuery(t *testing.T) {
	values, fieldNames, err := query.Values(new(OnlineTimeQuery))
	if err != nil {
		panic(err)
	}
	fmt.Println(query.Encode(values, '&', fieldNames))
}

func TestOrderQuery(t *testing.T) {
	values, fieldNames, err := query.Values(new(OrderQuery))
	if err != nil {
		panic(err)
	}
	fmt.Println(query.Encode(values, '&', fieldNames))
}
