package constants

type contextKey string

const (
	COCKROACH_DB = "cockroach"
	SQLITE_DB   = "sqlite"
)

var (
	DB_CONN_CTX_KEY = contextKey("dbConn")
)
