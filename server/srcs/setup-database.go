package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type db = pgxpool.Pool

func createTable(db *db, name string, fields []string) {
	var query = "CREATE TABLE " + name + " ("
	for i, str := range fields {
		query += str
		if i != len(fields) - 1 {
			query += ", "
		}
	}
	query += ");"
	_, err := db.Exec(context.Background(), query)
	if err == nil {
		log.Printf("created %s table", name)
	} else {
		log.Println(err.Error())
	}
}

func createUserTable(db *db) {
	createTable(db, "users", []string{
		"id int NOT NULL",
		"username text NOT NULL",
		"password text NOT NULL",
		"PRIMARY KEY (id)",
	})
}

func SetupDB() *db {
	var url = fmt.Sprintf("postgres://%s:%s@postgres:5432/postgres",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		panic(err)
	}
	createUserTable(db)
	return db
}
