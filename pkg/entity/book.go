package entity

type Book struct {
	ID            string
	Title         string
	AuthorID      string
	NumberOfPages int
	Description   string
}

func (b *Book) GetID() string {
	return b.ID
}
