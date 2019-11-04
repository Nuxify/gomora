package datetime

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
)

//NsTo24 - converts nanoseconds to 24 hour format hh, mm, and ss
func NsTo24(input string) (h int, m int, s int) {
	floatNanoseconds, _ := strconv.ParseFloat(input, 64)
	floatSeconds := floatNanoseconds / 1000000000

	hours := math.Floor(floatSeconds / 60 / 60)
	modSeconds := int(floatSeconds) % (60 * 60)
	minutes := math.Floor(float64(modSeconds) / 60)
	seconds := int(floatSeconds) % 60

	return int(hours), int(minutes), int(seconds)
}

//ParseLocal - converts string to scylla accepted time format
func ParseLocal(timeStr string) time.Time {
	t, _ := dateparse.ParseLocal(timeStr)

	return t
}

//TimeToManilaSeconds - converts given time (in UTC) to seconds in Asia/Manila
func TimeToManilaSeconds(t time.Time) int64 {
	return t.Unix() + 28800
}

//ToScyllaTimestamp - converts given string time with 23:59:59 format of scylla
func ToScyllaTimestamp(timeStr string) time.Time {
	t := ParseLocal(fmt.Sprintf("%s 23:59:59", timeStr))
	//t, _ := time.Parse("2006-01-02 15:04:05", timeStr) //! reference by me, @author: karl

	return t
}
