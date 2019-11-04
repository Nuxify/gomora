//TODO: string conversion helper
package string

import (
	"strconv"
)

func Parse(val interface{}) string {
	switch val.(type) {
	case float64:
		return strconv.FormatFloat(val.(float64), "f")
	}
}
