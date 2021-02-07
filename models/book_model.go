package models

type Book struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Isbn        string `json:"isbn"`
}

func (b *Book) TableName() string {
	return "book"
}
