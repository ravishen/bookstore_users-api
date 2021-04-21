package date_utils

import "time"

const (
	api_date_layout = "2006-01-02T15:04:05Z07:00"
)

func GetNowString() string {
	currTime := time.Now().UTC()
	return currTime.Format(api_date_layout)
}
