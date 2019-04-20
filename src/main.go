package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Rocko10/Go-TODO-list/src/handler/todos"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", todos.List)
	router.HandleFunc("/todos/new", todos.New).Methods("POST")
	router.HandleFunc("/todos/toggle", todos.Toggle).Methods("POST")
	router.HandleFunc("/todos/delete", todos.Delete).Methods("POST")

	http.ListenAndServe(":8080", router)
}
