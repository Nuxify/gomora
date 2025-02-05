package types

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// MySQLDBHandlerInterface contains the implementable methods for the MySQL DB handler
type MySQLDBHandlerInterface interface {
	// Begin starts a new transaction
	Begin() (*sqlx.Tx, error)
	// Execute executes the mysql statement following NamedExec
	Execute(stmt string, model interface{}) (sql.Result, error)
	// Query selects rows given by the sql statement
	Query(qstmt string, model interface{}, bindModel interface{}) error
	// QueryRow selects a row given by the sql statement
	QueryRow(qstmt string, model interface{}, bindModel interface{}) error
}
