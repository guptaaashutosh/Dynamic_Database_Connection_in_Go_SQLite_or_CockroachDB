package middlewares

import (
	"context"
	"database/sql"
	"net/http"
	"tutorials/dynamic_db_conn_project/constants"
)

func SetDBConnection(dbConn *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), constants.DB_CONN_CTX_KEY, dbConn)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
