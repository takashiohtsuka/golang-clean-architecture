package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS golang_clean_architecture_test CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_as_ci")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("golang_clean_architecture_test created successfully")
}
