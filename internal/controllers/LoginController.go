package controllers

import (
	"log"
	"myauth/internal/providers"
	"net/http"
)

type LoginController struct {
	Provider *providers.DatabaseProvider
}

type userModel struct {
	id       int
	username string
	password string
}

func (controller *LoginController) LoginHandler(w http.ResponseWriter, r *http.Request) {

	log.Default().Printf("Request received!!!\n")

	var params map[string]string = map[string]string{}
	params["username"] = "test"
	params["password"] = "test"

	rows := controller.Provider.Get("users", params)
	var user userModel = userModel{}

	if !rows.Next() {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	scanErr := rows.Scan(&user.id, &user.username, &user.password)

	if scanErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
