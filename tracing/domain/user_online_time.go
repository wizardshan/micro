package domain

import (
	"github.com/elliotchance/pie/v2"
	"time"
)

type UserOnlineTimes []*UserOnlineTime

func (dom UserOnlineTimes) Duration() int {
	sum := 0
	pie.Each(dom, func(item *UserOnlineTime) {
		sum += item.Duration
	})
	return sum
}

type UserOnlineTime struct {
	GameID    int
	ServerID  int
	Date      string
	StartTime time.Time
	EndTime   time.Time
	Duration  int
}
