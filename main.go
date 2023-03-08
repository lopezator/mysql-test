package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

func querySQLMode(dsn string) (string, error) {
	var sqlMode string
	if db, err := sql.Open("mysql", dsn); err == nil {
		row := db.QueryRowContext(context.Background(), "SELECT @@SESSION.sql_mode;")
		if err := row.Scan(&sqlMode); err != nil {
			return "", err
		}
	}
	return sqlMode, nil
}

func main() {
	// Set connection string.
	dsnStr := "root:root@tcp(localhost:3307)/test"

	// Initialize dsn object.
	var dsn *mysql.Config
	dsn, err := mysql.ParseDSN(dsnStr)
	if err != nil {
		log.Fatal(err)
	}
	dsn.Params = map[string]string{}

	// Get current sql_mode.
	dsn.Params["sql_mode"], err = querySQLMode(dsnStr)
	if err != nil {
		log.Fatal(err)
	}

	// Remove the NO_ZERO_DATE flag.
	dsn.Params["sql_mode"] = strings.Replace(dsn.Params["sql_mode"], "NO_ZERO_DATE,", "", -1)

	// Enclose the param in quotes.
	dsn.Params["sql_mode"] = fmt.Sprintf("%q", dsn.Params["sql_mode"])

	// Open connection.
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatalf("impossible open database: %s", err)
	}
	defer db.Close()

	// Insert the record.
	query := "INSERT INTO table_name (datetime) VALUES (?)"
	if _, err = db.ExecContext(context.Background(), query, time.Time{}); err != nil {
		log.Fatalf("impossible insert teacher: %s", err)
	}
}
