package snowflake

import (
	"fmt"
	"testing"
)

func TestGenerateSnowflake(t *testing.T) {
	id, _ := Generate()
	fmt.Println(id)
}
