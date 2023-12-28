package types

type ConnectionParams struct {
	Dial       string // register dial
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string
}

type SSHConnectionParams struct {
	SSHHost     string
	SSHPort     string
	SSHUsername string
	SSHPassword string
}
