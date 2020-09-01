package database

import (
	"database/sql"
	"log"
	_ "github.com/bmizerany/pq"

)

//FUnciones que empeizan con Mayusculas se exportan automaticamente
func GetConnection() *sql.DB {
	//connStr := "postgres://postgres:postgres@localhost/go_api_rest?sslmode=disable"
	db, err := sql.Open("postgres", "dbname=go_api_rest user=postgres password=postgres port=5432 sslmode=disable");
	if err != nil {
		log.Fatal(err)
	}

	return db
}