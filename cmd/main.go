package main

import (
	"log"
	"myauth/internal/controllers"
	"myauth/internal/providers"
	"net/http"
)

func main() {

	dbProvider, dbError := providers.InitDatabase()

	if dbError != nil {
		panic("database couldn't be initialized")
	}

	// Setting-up the controllers
	var loginController controllers.LoginController = controllers.LoginController{
		Provider: dbProvider,
	}

	var authController controllers.AuthController = controllers.AuthController{
		Provider: dbProvider,
	}

	// Creating the server instance and setting-up
	// the endpoints.
	// Also we map each endpoint to a controller.
	var addr string = ":27000"

	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", loginController.LoginHandler)
	mux.HandleFunc("POST /auth", authController.AuthHandler)

	log.Default().Printf("Server starting at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
	log.Default().Print("Server initialized!!!")
}
