package file

import "time"

func GetHourDiffer(start_time, end_time time.Time) int64 {
	var hour int64
	//	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	//	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if true == start_time.Before(end_time) {
		diff := end_time.Unix() - start_time.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}
