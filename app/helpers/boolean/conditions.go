package boolean

// IsSet - returns a boolean response if a map contains the certain key
func IsSet(list map[string]interface{}, key string) bool {
	_, found := list[key]
	if !found {
		return false
	}

	return true
}
