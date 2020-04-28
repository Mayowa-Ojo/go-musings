package database

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/go-musings/server/database/models"
	// import pg driver
	_ "github.com/lib/pq"
)

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

// QueryInsertRow - insert row into table
func QueryInsertRow(db *sql.DB, data models.Book) error {
	title := data.Title
	author := data.Author

	query := fmt.Sprintf(`INSERT INTO books (title, author, created_at) VALUES ('%s', '%s', NOW());`, title, author)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

// QueryRows - fetch all rows from table
func QueryRows(db *sql.DB) ([]models.Book, error) {
	var books []models.Book
	query := `SELECT * FROM books`

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b models.Book

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
func QueryRow(db *sql.DB, table string, id int) (models.Book, error) {
	var b models.Book

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

// QueryUpdateRow -
func QueryUpdateRow(db *sql.DB, id int, table string, update models.Book) error {
	query := `
		UPDATE %s
		SET title = '%s'
		WHERE id = %d;
	`
	q := fmt.Sprintf(query, table, update.Title, id)

	_, err := db.Exec(q)

	if err != nil {
		return err
	}

	return nil
}

// QueryDeleteRow -
func QueryDeleteRow(db *sql.DB, id int, table string) error {
	query := "DELETE FROM %s WHERE id = %d"
	q := fmt.Sprintf(query, table, id)

	_, err := db.Exec(q)

	if err != nil {
		return err
	}

	return nil
}
