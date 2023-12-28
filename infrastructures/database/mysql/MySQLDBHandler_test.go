package mysql

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"gomora/infrastructures/database/mysql/types"
)

func TestConnection(t *testing.T) {
	// load our environmental variables.
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error(err)
	}

	db := &MySQLDBHandler{}
	err := db.Connect(types.ConnectionParams{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		t.Error("connection error")
		return
	}

	t.Log("connection success")
}

func TestSSHConnection(t *testing.T) {
	// load our environmental variables.
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error(err)
	}

	db := &MySQLDBHandler{}
	err := db.ConnectViaSSH(types.SSHConnectionParams{
		SSHHost:     os.Getenv("DB_SSH_HOST"),
		SSHPort:     os.Getenv("DB_SSH_PORT"),
		SSHUsername: os.Getenv("DB_SSH_USER"),
		SSHPassword: os.Getenv("DB_SSH_PASSWORD"),
	}, types.ConnectionParams{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		t.Error("connection error via ssh")
		return
	}

	t.Log("connection success via ssh")
}
