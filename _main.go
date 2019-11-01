package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	log.Fatal(http.ListenAndServe(":12345", router))
}
