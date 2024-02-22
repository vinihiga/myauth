package controllers

import (
	"encoding/json"
	"log"
	"myauth/internal/providers"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type LoginController struct {
	Provider   *providers.DatabaseProvider
	PrivateKey []byte
}

type userModel struct {
	Id       int
	Username string
	Password string
}

// LoginHandler handles the login process for the LoginController.
// It takes in the http.ResponseWriter and http.Request as parameters.
// It does not return anything.
func (controller *LoginController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if controller.Provider == nil {
		log.Fatal("no provider was found!!!")
		return
	}

	log.Default().Printf("request received!!!\n")

	var user userModel = userModel{}

	_ = json.NewDecoder(r.Body).Decode(&user)

	var params map[string]string = map[string]string{}
	params["username"] = user.Username
	params["password"] = user.Password

	rows := controller.Provider.Get("users", params)

	if !rows.Next() {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	// I'm mocking here the private key.
	// Remember: This is a bad practice, but for testing
	// and learning purpose, I'm adding here.
	var token *jwt.Token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"foo": "bar",
		},
	)

	var result, tokenErr = token.SignedString(controller.PrivateKey)

	if tokenErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
