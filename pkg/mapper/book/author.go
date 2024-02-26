package book

import (
	"GO_RESTful_API/pkg/entities/impl/book"
	"encoding/json"
)

func MapJsonToAuthor(jsonData []byte) (book.Author, error) {
	var author book.Author
	err := json.Unmarshal(jsonData, &author)
	if err != nil {
		return book.Author{}, err
	}
	return author, nil
}
