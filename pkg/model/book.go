package model

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	NumberOfPages int    `json:"number-of-pages"`
	Description   string `json:"description"`
	AuthorID      string `json:"author-id"`
}
