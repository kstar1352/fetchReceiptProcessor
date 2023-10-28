package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receiptMap = make(map[string]Receipt)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

func ProcessReceipt(rw http.ResponseWriter, r *http.Request) {
	ok := true
	var id string
	for ok {
		id = uuid.NewString()
		_, ok = receiptMap[id]

	}
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		// respond somehow
		rw.Write([]byte(err.Error()))
	}
	receiptMap[id] = receipt
	resp := ProcessReceiptResponse{
		ID: id,
	}
	JsonWriter(rw, resp)
}

type PointsResp struct {
	Points int `json:"points"`
}

func GetPoints(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	receipt, ok := receiptMap[id]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("No Existing Receipt ID"))
	}
	points := CalcPoints(receipt)
	resp := PointsResp{
		Points: points,
	}
	JsonWriter(rw, resp)

}
