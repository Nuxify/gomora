package mysql

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	// load our environmental variables.
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error(err)
	}

	db := &MySQLDBHandler{}
	err := db.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	if err != nil {
		t.Error("connection error")
		return
	}

	t.Log("connection success")
}
