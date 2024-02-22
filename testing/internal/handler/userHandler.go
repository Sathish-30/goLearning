package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sathish-30/go-testing/internal/model"
	"github.com/sathish-30/go-testing/internal/util"
)

func UserHomeRoute(w http.ResponseWriter, request *http.Request) {
	userData := model.User{UserId: 1, UserName: "sathish", UserAge: 21}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(userData)
	util.CheckErr(err)
}
