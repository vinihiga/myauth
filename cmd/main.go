package main

import (
	"log"
	"myauth/internal/controllers"
	"net/http"
)

func main() {

	var addr string = ":27000"
	var loginController controllers.LoginController = controllers.LoginController{}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", loginController.LoginHandler)

	log.Default().Printf("Server starting at %s", addr)
	log.Fatal(http.ListenAndServe(":27000", mux))
	log.Default().Print("Server initialized!!!")
}
