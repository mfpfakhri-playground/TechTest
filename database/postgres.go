package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

// Initiation connect to database
func Initiation() {
	log.Println("database: trying connect to database")

	createConnection()
	log.Println("database: successfully connected to database")
}

// CreateConnection open new connection Postgres
func createConnection() error {

	dsn := "host=localhost port=5432 user=postgres password=standar123 dbname=tech_test sslmode=disable"

	dbCon, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("could not open postgres database dsn: %s\n", err.Error())
		return err
	}
	err = dbCon.Ping()
	if err != nil {
		log.Printf("could not ping postgres database: %s\n", err.Error())
		return err
	}
	log.Printf("database postgres: Connected!\n")
	dbCon.SetMaxOpenConns(100)
	dbCon.SetMaxIdleConns(10)
	dbCon.SetConnMaxLifetime(time.Duration(300 * time.Second))
	dbConn = dbCon
	return nil
}

// GetPostGresDBconn get pointer database connection
func GetPostGresDBconn() *sql.DB {
	return dbConn
}

// ReplaceDBConn this function for testing purpose only
// it will replace dbConn pointer to DBMock pointer from go-sqlmock
func ReplaceDBConn(DBMockConn *sql.DB) {
	dbConn = DBMockConn
}
