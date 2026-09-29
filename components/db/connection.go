package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(driver string, connection string) {
	var err error
	DB, err = sql.Open(driver, connection)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	//DB.SetConnMaxLifetime(time.Duration(5) * time.Second)
	//DB.SetMaxOpenConns(100)
}
