package handler

import (
	"errors"
	"net/http"

	"github.com/sathish-30/jsonApiProject/internal/model"
	"github.com/sathish-30/jsonApiProject/internal/util"
)

var Books = []model.Book{
	{BookId: 1, BookName: "Think like a monk", Amount: 2000},
	{BookId: 2, BookName: "Alchemist", Amount: 1000},
	{BookId: 3, BookName: "Atomic habits", Amount: 3000},
}

type ApiFuncCall func(w http.ResponseWriter, r *http.Request) error

func HomeRouteHandler(w http.ResponseWriter, r *http.Request) error {
	return util.WriteJson(w, http.StatusAccepted, map[string]string{
		"message": "In home route",
	})
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) error {
	return util.WriteJson(w, http.StatusAccepted, Books)
}

func AddBookToShop(w http.ResponseWriter, r *http.Request) error {
	payload, err := util.ReadJson(r)
	if err != nil {
		return errors.New("unable to parse request body to a struct")
	}
	book := model.Book(payload)
	Books = append(Books, book)
	return util.WriteJson(w, http.StatusAccepted, map[string]any{
		"Message": "Added book to the database",
		"books":   Books,
	})
}
