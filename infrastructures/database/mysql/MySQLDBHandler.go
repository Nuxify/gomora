package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"gomora/infrastructures/database/mysql/types"
)

// MySQLDBHandler handles mysql operations
type MySQLDBHandler struct {
	Conn *sqlx.DB
}

type viaSSHDialer struct {
	client *ssh.Client
}

// Connect opens a new connection to the mysql interface
func (h *MySQLDBHandler) Connect(params types.ConnectionParams) error {
	if len(params.Dial) == 0 {
		params.Dial = "tcp" // default
	}

	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true&sql_mode=TRADITIONAL", params.DBUsername, params.DBPassword, params.Dial, params.DBHost, params.DBPort, params.DBDatabase))
	if err != nil {
		return err
	}

	conn.SetConnMaxLifetime(time.Minute * 4)
	h.Conn = conn

	err = conn.Ping()
	if err != nil {
		connErr := fmt.Errorf("[SERVER] Error connecting to the database! %s", err.Error())

		return connErr
	}

	fmt.Println("[SERVER] Database connected successfully")

	return nil
}

// ConnectViaSSH opens a new connection to the mysql interface through ssh
// https://gist.github.com/vinzenz/d8e6834d9e25bbd422c14326f357cce0
// https://unix.stackexchange.com/a/415266
func (h *MySQLDBHandler) ConnectViaSSH(paramsSSH types.SSHConnectionParams, params types.ConnectionParams) error {
	var agentClient agent.Agent

	// establish a connection to the local ssh-agent
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return err
	}

	// create a new instance of the ssh agent
	agentClient = agent.NewClient(conn)

	// the client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User:            paramsSSH.SSHUsername,
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// when the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}

	// when there's a non empty password add the password AuthMethod
	if paramsSSH.SSHPassword != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return paramsSSH.SSHPassword, nil
		}))
	}

	// connect to the SSH Server
	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", paramsSSH.SSHHost, paramsSSH.SSHPort), sshConfig)
	if err != nil {
		return err
	}

	// now we register the ViaSSHDialer with the ssh connection as a parameter
	mysql.RegisterDialContext(params.Dial, func(_ context.Context, addr string) (net.Conn, error) {
		dialer := &viaSSHDialer{sshConn}
		return dialer.Dial(addr)
	})

	// connect to database
	err = h.Connect(params)
	if err != nil {
		return err
	}

	return nil
}

// Execute executes the mysql statement following NamedExec
// It requires a valid sql statement and its struct
func (h *MySQLDBHandler) Execute(stmt string, model interface{}) (sql.Result, error) {
	res, err := h.Conn.NamedExec(stmt, model)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query selects rows given by the sql statement
// It requires the statement, the model to bind the statement, and the target bind model for the results
func (h *MySQLDBHandler) Query(qstmt string, model interface{}, bindModel interface{}) error {
	nstmt, err := h.Conn.PrepareNamed(qstmt)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	err = nstmt.Select(bindModel, model)
	return err
}

// QueryRow selects a row given by the sql statement
// It requires the statement, the model to bind the statement, and the target bind model for the result
func (h *MySQLDBHandler) QueryRow(qstmt string, model interface{}, bindModel interface{}) error {
	nstmt, err := h.Conn.PrepareNamed(qstmt)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	err = nstmt.Get(bindModel, model)
	return err
}

func (v *viaSSHDialer) Dial(addr string) (net.Conn, error) {
	return v.client.Dial("tcp", addr)
}
