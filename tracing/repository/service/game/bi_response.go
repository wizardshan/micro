package game

import (
	"tracing/domain"
	"tracing/pkg/timeformat"
)

type UserOnlineTimes []*UserOnlineTime

type UserOnlineTime struct {
	GameID    int                 `json:"game_id,string"`
	ServerID  int                 `json:"server_id,string"`
	Date      string              `json:"date"`
	StartTime timeformat.DateTime `json:"start_time"`
	EndTime   timeformat.DateTime `json:"end_time"`
	Duration  int                 `json:"online_time,string"`
}

func (entity *UserOnlineTime) Mapper() *domain.UserOnlineTime {
	if entity == nil {
		return nil
	}
	dom := new(domain.UserOnlineTime)
	dom.GameID = entity.GameID
	dom.ServerID = entity.ServerID
	dom.Date = entity.Date
	dom.StartTime = entity.StartTime.Time()
	dom.EndTime = entity.EndTime.Time()
	dom.Duration = entity.Duration

	return dom
}

func (entities UserOnlineTimes) Mapper() domain.UserOnlineTimes {
	var doms domain.UserOnlineTimes
	for _, entity := range entities {
		doms = append(doms, entity.Mapper())
	}
	return doms
}
