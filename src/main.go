package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Rocko10/todoApp/src/handler/todos"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", todos.List)
	router.HandleFunc("/todos/new", todos.New).Methods("POST")

	http.ListenAndServe(":8080", router)
}
