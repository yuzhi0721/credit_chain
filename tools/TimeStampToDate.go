package tools

import "time"

func TimeStampToDate(t int64)string{
	var date string
	if t != 0{
		date = time.Unix(t,0).Format("2006-01-02")
	}
	return date
}