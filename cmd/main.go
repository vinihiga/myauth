package main

import (
	"log"
	"myauth/internal/controllers"
	"myauth/internal/providers"
	"net/http"
	"os"
)

func readPrivateKey() []byte {
	content, readErr := os.ReadFile("env/mock_private_key.txt")

	if readErr != nil {
		return []byte{}
	}

	return content
}

func main() {

	// Setting-up the private key.
	// In this case, we are using a mock key.
	// Real environments shouldn't commit the key.
	var privateKey = readPrivateKey()

	// Initializing database provider.
	// Here, basically, we start communicating with our
	// PostgreSQL database.
	dbProvider, dbError := providers.InitDatabase()

	if dbError != nil {
		panic("database couldn't be initialized")
	}

	// Setting-up the controllers.
	// We are going to have 2 endpoints.
	// 1st - Login Controller.
	// 2nd - Auth Controller.
	var loginController controllers.LoginController = controllers.LoginController{
		Provider:   dbProvider,
		PrivateKey: privateKey,
	}

	var authController controllers.AuthController = controllers.AuthController{
		Provider:   dbProvider,
		PrivateKey: privateKey,
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
