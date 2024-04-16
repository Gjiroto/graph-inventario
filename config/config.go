package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:Ha3223730151@tcp(localhost:3306)/inventory")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection successful")
}
