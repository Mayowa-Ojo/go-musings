package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// import pg driver
	_ "github.com/lib/pq"
)

// Service - holds a DB type
type Service struct {
	db *sql.DB
}

// Book - describes a book structure
type Book struct {
	ID        int
	Title     string
	Author    string
	CreatedAt time.Time
}

/*
func main() {
	// s := service{}
	connStr := "postgres://postgres:adebayor@localhost/test?sslmode=disable"

	db, err := connectDB("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// rows, err := db.Query("SELECT * FROM books")
	// books, err := queryRows(db)
	// -------------------------------------------------
	// data := book{}
	// data.title = `The Subtle Art of Not Giving a F*ck`
	// data.author = `Mark Manson`
	// err = insertRow(db, data)
	// -------------------------------------------------
	err = createTable(db, "authors")
	// -------------------------------------------------
	book, err := queryRow(db, "books", 4)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("book: %v\n", book.createdAt)
}
*/

// ConnectDB - connect to DB
func ConnectDB(dbType, conn string) (*sql.DB, error) {
	db, err := sql.Open(dbType, conn)

	if err != nil {
		return nil, err
	}

	fmt.Println("[x] --psql: database connected...")

	return db, nil
}

// CreateTable - create a table in DB
func CreateTable(db *sql.DB, table string) error {
	query := `
		CREATE TABLE %s (
			id serial PRIMARY KEY,
			title VARCHAR (150) NOT NULL,
			author VARCHAR (100) NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
	`

	q := fmt.Sprintf(query, table)

	_, err := db.Exec(q)

	if err != nil {
		return err
	}

	return nil
}

// InsertRow - insert row into table
func InsertRow(db *sql.DB, data Book) error {
	title := data.Title
	author := data.Author

	query := fmt.Sprintf(`INSERT INTO books (title, author, created_at) VALUES ('%s', '%s', NOW());`, title, author)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	fmt.Println("row created...")
	return nil
}

// QueryRows - fetch all rows from table
func QueryRows(db *sql.DB) ([]Book, error) {
	var books []Book
	query := `SELECT * FROM books`

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b Book

		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt)

		if err != nil {
			log.Fatal(err)
			break
		}

		books = append(books, b)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return books, nil
}

// QueryRow - fetch a single row from table
func QueryRow(db *sql.DB, table string, id int) (Book, error) {
	// var (
	// 	title     string
	// 	author    string
	// 	createdAt time.Time
	// )

	var b Book

	query := `SELECT * FROM %s WHERE id = %d`
	q := fmt.Sprintf(query, table, id) // format query string

	row, err := db.Query(q)

	if err != nil {
		return b, err
	}

	_ = row.Next()
	err = row.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt)

	if err != nil {
		return b, err
	}

	return b, nil
}
