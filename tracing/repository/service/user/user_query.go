package user

import (
	"fmt"
	"time"
)

type Query struct {
	AgentType int    `url:"agenttype"`
	Field     string `url:"field"`
	Timestamp int64  `url:"timestamp"`
	Source    string `url:"source"`
}

func defaultQuery() Query {
	return Query{
		AgentType: 39,
		Field:     "userinfo,qiyi_vip_info",
		Timestamp: time.Now().UnixMicro(),
		Source:    "youxi",
	}
}

type QueryUser struct {
	Query
}

type QueryUsers struct {
	Query
	UserIDs string `url:"uids"`
}

func (q *QueryUsers) Encode() string {
	return fmt.Sprintf("agenttype=%d|field=%s|source=%s|timestamp=%d|uids=%s", q.AgentType, q.Field, q.Source, q.Timestamp, q.UserIDs)
}

func (q *QueryUsers) Build() string {
	return fmt.Sprintf("agenttype=%d&field=%s&source=%s&timestamp=%d&uids=%s", q.AgentType, q.Field, q.Source, q.Timestamp, q.UserIDs)
}
