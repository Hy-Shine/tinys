package go_utils

import "time"

// TimeToStr convert time to yyyy-mm-dd hh24:mi:ss string
func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// StrParseToTime return time of string
func StrParseToTime(t string, location ...interface{}) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	if len(location) > 0 {
		local := time.Now().Location()
		return time.ParseInLocation(layout, t, local)
	}

	return time.Parse(layout, t)
}
