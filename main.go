package main

import (
	"log"
	"net/http"
	"receiptProcessor/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", server.ProcessReceipt).Methods(http.MethodPost)
	router.HandleFunc("/receipts/{id}/points", server.GetPoints).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":3000", router))
}
