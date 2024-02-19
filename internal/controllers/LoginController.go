package controllers

import "net/http"

type LoginController struct {
}

func (LoginController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
