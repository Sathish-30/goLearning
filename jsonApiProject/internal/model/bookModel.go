package model

type Book struct {
	BookId   int    `json:"bookId"`
	BookName string `json:"bookName"`
	Amount   int    `json:"amount"`
}
