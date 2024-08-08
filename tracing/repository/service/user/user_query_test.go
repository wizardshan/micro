package user

import (
	"testing"
	"tracing/pkg/query"
)

func TestQueryUserEncode(t *testing.T) {

	q, fieldNames, err := query.Values(new(QueryUser))
	if err != nil {
		panic(err)
	}
	//fmt.Println(query.Encode(q, '|', []string{"sign"}, fieldNames))
	//fmt.Println(query.Encode(q, '&', []string{}, fieldNames))
}
