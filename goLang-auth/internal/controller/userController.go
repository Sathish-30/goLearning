package controller

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sathish-30/golang-auth/internal/auth"
	"github.com/sathish-30/golang-auth/internal/database"
	"github.com/sathish-30/golang-auth/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(models.Error{
			ErrorMessage: "Bad Request",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		json.NewEncoder(w).Encode(models.Error{
			ErrorMessage: "Failed to hash password",
		})
	}

	newUser := models.User{
		Email:    user.Email,
		Password: string(hash),
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		json.NewEncoder(w).Encode(models.Error{
			ErrorMessage: "Failed to create user",
		})
	}

	json.NewEncoder(w).Encode(&newUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.User
	w.Header().Set("Content-Type", "application/json")
	// Fetching data from request body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		json.NewEncoder(w).Encode(models.Error{
			ErrorMessage: "Unable to fetch request body data",
		})
		return
	}

	var user models.User
	// Fetching record from Db with the help of email Id
	database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		json.NewEncoder(w).Encode(models.Error{
			ErrorMessage: "Invalid Email",
		})
		return
	}

	// Checking whether it is a valid password or not
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		json.NewEncoder(w).Encode(models.Error{
			ErrorMessage: "Invalid password",
		})
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5).Unix()

	standardClaims := jwt.StandardClaims{
		ExpiresAt: expirationTime,
	}

	claims := &auth.Claims{
		Name:           user.Email,
		StandardClaims: standardClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(models.Error{
			ErrorMessage: "Not able to generate token",
		}); err != nil {
			log.Println(err)
		}
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Minute * 5),
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	if err := json.NewEncoder(w).Encode(map[string]string{
		"Token": tokenString,
	}); err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func Home(w http.ResponseWriter, r *http.Request) {

}
