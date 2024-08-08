package bi

import (
	"fmt"
	"testing"
	"tracing/pkg/query"
)

func TestQueryUserEncode(t *testing.T) {
	values, fieldNames, err := query.Values(new(OnlineTimeQuery))
	if err != nil {
		panic(err)
	}
	fmt.Println(query.Encode(values, '&', fieldNames))
}
