package structs

import (
	"github.com/kabaluyot/structs"
)

// StructToMap - converts given struct to map type
func StructToMap(item interface{}) map[string]interface{} {
	return structs.Map(item)
}
