package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connection() (*sqlx.DB){
	godotenv.Load(".env")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", os.Getenv("USER"),os.Getenv("DB_NAME"),os.Getenv("PASSWORD"),os.Getenv("HOST")))
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db;

}

func Query(query string) *sqlx.Rows{
	db := connection()

	defer db.Close()

	rows, err := db.Queryx(query)

	if err != nil {
        panic(err.Error())
    }

	return rows;
}

func QueryRow(query string, args ...interface{} ) *sqlx.Row{
	db := connection()

	defer db.Close()

	row := db.QueryRowx(query, args)

	if row != nil {
		panic(row)
    }

	return row;
}