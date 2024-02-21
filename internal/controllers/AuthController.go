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
	Provider *providers.DatabaseProvider
}

type requestBodyMock struct {
	Foo string
}

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

	token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("seems like the signing method is different that HMAC")
		// }

		return "test123", nil
	})

	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	// if claimErr != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte{})
	// 	return
	// }

	var response requestBodyMock = requestBodyMock{}
	response.Foo = fmt.Sprint(claims["foo"])

	jsonContent, parseErr := json.Marshal(response)

	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}
