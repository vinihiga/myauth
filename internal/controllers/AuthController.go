package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"myauth/internal/providers"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	Provider   *providers.DatabaseProvider
	PrivateKey []byte
}

type requestBodyMock struct {
	Foo string
}

// AuthHandler is a handler for authentication.
//
// It takes a http.ResponseWriter and a http.Request as parameters.
func (controller *AuthController) AuthHandler(w http.ResponseWriter, r *http.Request) {
	if controller.Provider == nil {
		log.Fatal("no provider was found!!!")
		return
	}

	log.Default().Printf("request received!!!\n")

	buffer, readErr := io.ReadAll(r.Body)

	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
	}

	var tokenString string = strings.TrimSpace(string(buffer))

	token, parseErr := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return controller.PrivateKey, nil
	})

	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	var response requestBodyMock = requestBodyMock{}
	response.Foo = fmt.Sprint(claims["foo"])

	jsonContent, parseErr := json.Marshal(response)

	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}
