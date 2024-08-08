package bi

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"time"
	"tracing/pkg/encrypt"
)

type OnlineTimeQuery struct {
	encrypt.SignField
	UserID    int       `url:"uid"`
	Terminal  []int     `url:"terminal,comma"`
	StartDate time.Time `url:"start_date"`
	EndDate   time.Time `url:"end_date"`
}

func (q *OnlineTimeQuery) Encode() string {
	return fmt.Sprintf("end_date=%s&sign_key=%s&source=%d&start_date=%s&terminal=%s&uid=%d", q.EndDate.Format(time.DateOnly), q.SignKey, q.Source, q.StartDate.Format(time.DateOnly), pie.Join(q.Terminal, ","), q.UserID)
}

func (q *OnlineTimeQuery) String() string {
	return fmt.Sprintf("end_date=%s&source=%d&start_date=%s&terminal=%s&uid=%d", q.EndDate.Format(time.DateOnly), q.Source, q.StartDate.Format(time.DateOnly), pie.Join(q.Terminal, ","), q.UserID)
}
