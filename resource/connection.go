package resource

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"tutorials/dynamic_db_conn_project/config"
	"tutorials/dynamic_db_conn_project/constants"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once
var dbConn *sql.DB

func Connect() *sql.DB {
	once.Do(func() {
		Conf, err := config.Load()
		if err != nil {
			log.Fatal(err)
		}

		//dynamic db connection
		//cockroach db connection
		if Conf.DB_CONNECTION == constants.COCKROACH_DB {
			dbConn, err = sql.Open("postgres", Conf.COCKROACH_DB_URL)
			if err != nil {
				log.Fatal(err)
			}
		} else if Conf.DB_CONNECTION == constants.SQLITE_DB {
			//sqlite db connection
			dbConn, err = sqliteConnection()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("Invalid DB Connection")
		}

		err = dbConn.Ping()
		if err != nil {
			log.Fatal(err)
		}

	})
	return dbConn
}

func sqliteConnection() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	createSqliteTable(db)

	return db, err
}

func createSqliteTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT);
	DELETE FROM users;
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	fmt.Println("SQLite table created")
}
