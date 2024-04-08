package timex

import "time"

const (
	DateTimeLayout = "2006-01-02 15:04:05" // time.DateTime
	DateLayout     = "2006-01-02"          // time.DateOnly
	TimeLayout     = "15:04:05"            // time.TimeOnly
)

// TimeFormat convert time to yyyy-mm-dd hh24:mi:ss string
func TimeFormat(t time.Time) string {
	return t.Format(DateTimeLayout)
}

// StrParseToTime return time of string
func StrParseToTime(t string, location ...time.Location) (time.Time, error) {
	layout := DateTimeLayout
	if len(location) > 0 {
		local := time.Now().Location()
		return time.ParseInLocation(layout, t, local)
	}

	return time.Parse(layout, t)
}

func FirstDayOfMonth() {}

func FistDayOfCurrent()
