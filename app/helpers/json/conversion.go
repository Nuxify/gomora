package json

import (
	"encoding/json"
)

// MapToJSON - converts given map type to json string equivalent
func MapToJSON(val interface{}) string {
	res, _ := json.Marshal(val)

	return string(res)
}

// UnmarshalToMap - converts given json string to map[string]interface using Unmarshal
func UnmarshalToMap(val string) map[string]interface{} {
	// encode JSON string to JSON object
	parsedBytes := []byte(val)
	var data map[string]interface{}

	if err := json.Unmarshal(parsedBytes, &data); err != nil {
		data = nil
	}

	return data
}
