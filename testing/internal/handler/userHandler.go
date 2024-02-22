package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sathish-30/go-testing/internal/model"
	"github.com/sathish-30/go-testing/internal/util"
)

var users = []model.User{
	{1, "sathish", 21},
	{2, "ammu", 21},
	{3, "kaarthika", 19},
	{4, "rowdy", 22},
	{5, "dexter", 20},
}

func UserHomeRoute(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	util.CheckErr(err)
}

func HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Path[len("/user/"):], 10, 64)
	util.CheckErr(err)
	var result model.User
	for _, user := range users {
		if user.UserId == id {
			result = user
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	error2 := json.NewEncoder(w).Encode(result)
	util.CheckErr(error2)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	log.Println(newUser)
	util.CheckErr(err)
	users = append(users, newUser)
	log.Println("Added user")
}

func DeleteUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Path[len("/deleteUser/"):], 10, 64)
	util.CheckErr(err)
	var index int
	for i, user := range users {
		if user.UserId == id {
			index = i
		}
	}
	users = append(users[:index], users[index+1:]...)
	w.Header().Set("Content-Type", "application/json")
	error2 := json.NewEncoder(w).Encode(users)
	util.CheckErr(error2)
}
