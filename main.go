package main

import (
	app "contactsBook/authentication"
	"contactsBook/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/user/login", controllers.LoginAccount).Methods("POST")
	router.HandleFunc("/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/delete", controllers.DeleteContact).Methods("DELETE")
	router.HandleFunc("/contacts/update", controllers.UpdateContact).Methods("PUT")
	router.HandleFunc("/me/contacts", controllers.GetContacts).Methods("GET")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
