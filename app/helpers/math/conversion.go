package math

import (
	"strconv"
)

// ToFloat32 converts an input string to float32 type
func ToFloat32(num string) float32 {
	res, _ := strconv.ParseFloat(num, 32)

	return float32(res)
}

// ToFloat64 converts an input string to float64 type
func ToFloat64(num string) float64 {
	res, _ := strconv.ParseFloat(num, 64)

	return res
}

// ToInt32 converts an input string to int32 type
func ToInt32(num string) int32 {
	res, _ := strconv.ParseInt(num, 10, 32)

	return int32(res)
}

// ToInt64 converts an input string to int64 type
func ToInt64(num string) int64 {
	res, _ := strconv.ParseInt(num, 10, 64)

	return res
}

// ToUint32 converts an input string to uint32 type
func ToUint32(num string) uint32 {
	res, _ := strconv.ParseUint(num, 10, 32)

	return uint32(res)
}

// ToUint64 converts an input string to uint64
func ToUint64(num string) uint64 {
	res, _ := strconv.ParseUint(num, 10, 64)

	return res
}
