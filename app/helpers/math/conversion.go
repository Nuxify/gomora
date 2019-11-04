package math

import (
	"strconv"
)

func ToFloat32(num string) float32 {
	res, _ := strconv.ParseFloat(num, 32)

	return float32(res)
}

func ToFloat64(num string) float64 {
	res, _ := strconv.ParseFloat(num, 64)

	return res
}

func ToInt32(num string) int32 {
	res, _ := strconv.ParseInt(num, 10, 32)

	return int32(res)
}

func ToInt64(num string) int64 {
	res, _ := strconv.ParseInt(num, 10, 64)

	return res
}

func ToUint32(num string) uint32 {
	res, _ := strconv.ParseUint(num, 10, 32)

	return uint32(res)
}

func ToUint64(num string) uint64 {
	res, _ := strconv.ParseUint(num, 10, 64)

	return res
}
