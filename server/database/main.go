package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// import pg driver
	_ "github.com/lib/pq"
)

// DB -
type service struct {
	db *sql.DB
}

type book struct {
	id        int
	title     string
	author    string
	createdAt time.Time
}

func main() {
	// s := service{}
	connStr := "postgres://postgres:adebayor@localhost/test?sslmode=disable"

	db, err := connectDB("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// rows, err := db.Query("SELECT * FROM books")
	// books, err := queryRows(db)
	data := book{}
	data.title = `The Subtle Art of Not Giving a F*ck`
	data.author = `Mark Manson`
	err = insertRow(db, data)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("books: %v\n", books)
}

func connectDB(dbType, conn string) (*sql.DB, error) {
	db, err := sql.Open(dbType, conn)

	if err != nil {
		return nil, err
	}

	fmt.Println("[x] --psql: database connected...")

	return db, nil
}

func (s *service) createTable() error {
	db := s.db

	query := `
		CREATE TABLE books (
			id serial PRIMARY KEY,
			title VARCHAR (150) NOT NULL,
			author VARCHAR (100) NOT NULL,
			created_at TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func insertRow(db *sql.DB, data book) error {
	title := data.title
	author := data.author

	query := fmt.Sprintf(`INSERT INTO books (title, author, created_at) VALUES ('%s', '%s', NOW());`, title, author)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	fmt.Println("row created...")
	return nil
}

func queryRows(db *sql.DB) ([]book, error) {
	var books []book
	query := `SELECT * FROM books`

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b book

		err := rows.Scan(&b.id, &b.title, &b.author, &b.createdAt)

		if err != nil {
			log.Fatal(err)
			break
		}

		books = append(books, b)
	}

	err = rows.Err()

	if err != nil {

	}

	return books, nil
}

func queryRow(db *sql.DB) (book, error) {
	// var (
	// 	title     string
	// 	author    string
	// 	createdAt time.Time
	// )

	return book{}, nil
}
