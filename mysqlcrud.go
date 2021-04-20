package main

import (
	"log"
  "github.com/gorilla/mux"
  "net/http"
  "github.com/gosample/getrecord"
  "github.com/gosample/createrecord"
  "github.com/gosample/updaterecord"
  "github.com/gosample/deleterecord"
)

func main() {
	router := mux.NewRouter()
  router.HandleFunc("/", getrecord.FetchAllItems).Methods("GET")
  router.HandleFunc("/create", createrecord.CreateTodo).Methods("POST")
  router.HandleFunc("/update", updaterecord.UpdateTodo).Methods("PUT")
	router.HandleFunc("/delete", deleterecord.DeleteTodo).Methods("DELETE")

	http.Handle("/", router)
	log.Print("localhost:8080")
	http.ListenAndServe(":8080", nil)
}
