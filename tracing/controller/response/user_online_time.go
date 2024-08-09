package response

import (
	"github.com/elliotchance/pie/v2"
	"time"
	"tracing/domain"
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

func MapperUserOnlineTime(domUserOnlineTime *domain.UserOnlineTime) *UserOnlineTime {
	if domUserOnlineTime == nil {
		return nil
	}
	respUserOnlineTime := new(UserOnlineTime)
	respUserOnlineTime.GameID = domUserOnlineTime.GameID
	respUserOnlineTime.ServerID = domUserOnlineTime.ServerID
	respUserOnlineTime.Date = domUserOnlineTime.Date
	respUserOnlineTime.StartTime = domUserOnlineTime.StartTime
	respUserOnlineTime.EndTime = domUserOnlineTime.EndTime
	respUserOnlineTime.Duration = domUserOnlineTime.Duration
	return respUserOnlineTime
}

func MapperUserOnlineTimes(domUserOnlineTimes domain.UserOnlineTimes) UserOnlineTimes {
	return mapper(domUserOnlineTimes, MapperUserOnlineTime)
}
