package database

import "time"

// Book - describes a book structure
type Book struct {
	ID        int
	Title     string
	Author    string
	CreatedAt time.Time
}
