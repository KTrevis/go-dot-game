package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBPool = pgxpool.Pool
type DB = pgxpool.Conn

func createTable(db *DB, name string, fields []string) {
	var query = "CREATE TABLE " + name + " ("

	for i, str := range fields {
		query += str
		if i != len(fields) - 1 {
			query += ", "
		}
	}
	query += ");"

	_, err := db.Exec(context.TODO(), query)

	if err == nil {
		log.Printf("TABLE CREATED: %s", name)
	} else {
		log.Println(err.Error())
	}
}

func createUserTable(db *DB) {
	createTable(db, "users", []string{
		"id SERIAL PRIMARY KEY",
		"username TEXT NOT NULL UNIQUE",
		"password TEXT NOT NULL",
	})
}

func createCharactersTable(db *DB) {
	createTable(db, "characters", []string{
		"id SERIAL PRIMARY KEY",
		"user_id INTEGER REFERENCES users(id) ON DELETE CASCADE",
		"name TEXT NOT NULL UNIQUE",
		"class TEXT NOT NULL",
		"level INTEGER NOT NULL",
		"xp INTEGER NOT NULL",
		"x INTEGER NOT NULL DEFAULT 0",
		"y INTEGER NOT NULL DEFAULT 0",
	})
}

func createChunksTable(db *DB) {
	createTable(db, "maps", []string{
		"id SERIAL PRIMARY KEY",
		"name TEXT NOT NULL",
		"tiles TEXT NOT NULL",
		"x INTEGER NOT NULL",
		"y INTEGER NOT NULL",
		"UNIQUE (x, y)",
	})
}

func connectToDB() *DBPool {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	const format = "postgres://%s:%s@postgres:5432/postgres"
	url := fmt.Sprintf(format, user, password)
	db, err := pgxpool.New(context.TODO(), url)

	if err != nil {
		panic(err)
	}
	return db
}

func SetupDB() *DBPool {
	db := connectToDB()
	conn, _ := db.Acquire(context.TODO())
	defer conn.Release()

	createUserTable(conn)
	createCharactersTable(conn)
	createChunksTable(conn)

	return db
}
