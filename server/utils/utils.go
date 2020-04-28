package utils

import (
	"strconv"

	models "github.com/go-musings/server/database/models"
)

// FetchBook - retrieves book and index for given id
func FetchBook(books []models.Book, id string) (b models.Book, index *int) {

	for i, v := range books {
		if idInt, _ := strconv.ParseInt(id, 10, 32); v.ID == int(idInt) {
			b = books[i]
			index = &i
			break
		}
	}

	return b, index
}
