package timepoint

import (
	"github.com/golang-module/carbon/v2"
	"time"
)

func init() {
	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.Shanghai,
		WeekStartsAt: carbon.Sunday,
		Locale:       "zh-CN",
	})
}

func StartOfDay() time.Time {
	return carbon.Now().StartOfDay().StdTime()
}

func EndOfDay() time.Time {
	return carbon.Now().EndOfDay().StdTime()
}

func StartOfYesterday() time.Time {
	return carbon.Yesterday().StartOfDay().StdTime()
}

func EndOfYesterday() time.Time {
	return carbon.Yesterday().EndOfDay().StdTime()
}

func StartOfMonth() time.Time {
	return carbon.Now().StartOfMonth().StdTime()
}

func EndOfMonth() time.Time {
	return carbon.Now().EndOfMonth().StdTime()
}

func StartOfLastMonth() time.Time {
	return carbon.Now().SubMonth().StartOfMonth().StdTime()
}

func EndOfLastMonth() time.Time {
	return carbon.Now().SubMonth().EndOfMonth().StdTime()
}

func StartOfWeek() time.Time {
	return carbon.Now().StartOfWeek().StdTime()
}

func EndOfWeek() time.Time {
	return carbon.Now().EndOfWeek().StdTime()
}

func StartOfLastWeek() time.Time {
	return carbon.Now().SubWeek().StartOfWeek().StdTime()
}

func EndOfLastWeek() time.Time {
	return carbon.Now().SubWeek().EndOfWeek().StdTime()
}
