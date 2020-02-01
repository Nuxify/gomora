package arrays

import (
	"gopkg.in/inf.v0"
	"strings"
)

// Explode - equvalent to PHP explode
func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	}

	return strings.Split(text, delimiter)
}

// IsDecimalInArray - checks if a decimal item is found in array string
func IsDecimalInArray(word *inf.Dec, slice []*inf.Dec) bool {
	for _, item := range slice {
		if item.Cmp(word) == 0 {
			return true
		}
	}

	return false
}

// IsInt32InArray - checks if a Int32 item is found in array string
func IsInt32InArray(word int32, slice []int32) bool {
	for _, item := range slice {
		if item == word {
			return true
		}
	}

	return false
}

// IsStringInArray - checks if a string item is found in array string
func IsStringInArray(word string, slice []string) bool {
	for _, item := range slice {
		if item == word {
			return true
		}
	}

	return false
}

// MinMaxDecimal - find max and min value of decimal data type array
func MinMaxDecimal(array []*inf.Dec) (*inf.Dec, *inf.Dec) {
	max := array[0]
	min := array[0]

	for _, value := range array {
		if value.Cmp(min) == -1 {
			min = value
		}
		if value.Cmp(max) == 1 {
			max = value
		}
	}

	return min, max
}

// PrependMap - prepend a map item to exiting map array using iteration algorithm
func PrependMap(items []map[string]interface{}, newItem map[string]interface{}) []map[string]interface{} {
	var prependedArr []map[string]interface{}

	//append the new item
	prependedArr = append(prependedArr, newItem)

	//iterate the original array and add to prepended array
	for _, item := range items {
		prependedArr = append(prependedArr, item)
	}

	return prependedArr
}
