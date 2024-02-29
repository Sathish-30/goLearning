package handler

import (
	"errors"
	"net/http"
)

type ApiFuncCall func(w http.ResponseWriter, r *http.Request) error

func HomeRouteHandler(w http.ResponseWriter, r *http.Request) error {
	// return util.WriteJson(w, http.StatusAccepted, map[string]string{
	// 	"message": "In home route",
	// })
	return errors.New("new error occured")
}
