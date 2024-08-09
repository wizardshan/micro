package timeformat

import "time"

type DateTime string

func (dt DateTime) Time() time.Time {
	date, _ := time.Parse(time.DateTime, string(dt))
	return date
}
