package database

var connection string

func init() {
	connection = "SQL SERVER"
}

func GetConnection() string {
	return connection
}
