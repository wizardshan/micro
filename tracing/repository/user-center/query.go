package usercenter

import (
	"fmt"
	"time"
)

type Query struct {
	AgentType int    `url:"agenttype"`
	Field     string `url:"field"`
	Timestamp int64  `url:"timestamp"`
	Source    string `url:"source"`
	Sign      string `url:"sign"`
}

type QueryUser struct {
	Query
	UserIDs string `url:"uids"`
}

func (q *QueryUser) Encode() string {
	return fmt.Sprintf("agenttype=%d|field=%s|source=%s|timestamp=%d|uids=%s", q.AgentType, q.Field, q.Source, q.Timestamp, q.UserIDs)
}

func (q *QueryUser) Build() string {
	return fmt.Sprintf("agenttype=%d&field=%s&sign=%s&source=%s&timestamp=%d&uids=%s", q.AgentType, q.Field, q.Sign, q.Source, q.Timestamp, q.UserIDs)
}

type QueryUsers struct {
	Query
}

func defaultQuery() Query {
	return Query{
		AgentType: 39,
		Field:     "userinfo,qiyi_vip_info",
		Timestamp: time.Now().UnixMicro(),
		Source:    "youxi",
	}
}
