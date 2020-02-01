package snowflake

import (
	"os"
	"strconv"

	"github.com/kabaluyot/snowflake"
)

// Generate - generate a new snowflake ID.
func Generate() (string, error) {
	// Create a new Node with a Node number of 1
	intNode, _ := strconv.ParseInt(os.Getenv("SNOWFLAKE_NODE"), 10, 16)
	node, err := snowflake.NewNode(intNode)
	if err != nil {
		return "", err
	}

	// Generate a snowflake ID.
	id := node.Generate()

	return id.String(), nil
}
