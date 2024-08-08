package bi

import (
	"time"
	"tracing/domain"
)

type DateTime string

func (dt DateTime) Time() time.Time {
	date, _ := time.Parse(time.DateTime, string(dt))
	return date
}

type Response struct {
	Code int     `json:"code"`
	Msg  *string `json:"msg"`
	Data any     `json:"data"`
}

func (resp *Response) Success() bool {
	return resp.Code == 200
}

type UserOnlineTimes []*UserOnlineTime

//type UserOnlineTime struct {
//	GameID    int       `json:"game_id"`
//	ServerID  int       `json:"server_id"`
//	Date      string    `json:"date"`
//	StartTime time.Time `json:"start_time"`
//	EndTime   time.Time `json:"end_time"`
//	Duration  int       `json:"online_time"`
//}

type UserOnlineTime struct {
	GameID    int      `json:"game_id,string"`
	ServerID  int      `json:"server_id,string"`
	Date      string   `json:"date"`
	StartTime DateTime `json:"start_time"`
	EndTime   DateTime `json:"end_time"`
	Duration  int      `json:"online_time,string"`
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
