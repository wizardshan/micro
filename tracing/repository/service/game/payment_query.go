package game

import (
	"fmt"
	"time"
	"tracing/domain/money"
)

const (
	OrderMoneyTypeDaySum = iota + 1
	OrderMoneyTypeDayList
	OrderMoneyTypeDurationSum
	OrderMoneyTypeDurationList
)

type OrderMoneyFen int

type OrderQuery struct {
	SignField
	UserID    int        `url:"user_id"`
	MinMoney  money.Unit `url:"min_money"`
	MoneyType int        `url:"money_type"`
	StartTime time.Time  `url:"start_time"`
	EndTime   time.Time  `url:"end_time"`
}

func (q *OrderQuery) Encode() string {
	return fmt.Sprintf("end_time=%s&min_money=%d&money_type=%d&sign_key=%s&source=%d&start_time=%s&user_id=%d", q.EndTime.Format(time.DateOnly), q.MinMoney, q.MoneyType, q.SignKey, q.Source, q.StartTime.Format(time.DateOnly), q.UserID)
}

func (q *OrderQuery) String() string {
	return fmt.Sprintf("end_time=%s&min_money=%d&money_type=%d&source=%d&start_time=%s&user_id=%d", q.EndTime.Format(time.DateOnly), q.MinMoney, q.MoneyType, q.Source, q.StartTime.Format(time.DateOnly), q.UserID)
}
