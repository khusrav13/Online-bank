package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Transaction struct {
	ID                    int64  `json:"id"`
	Date                  string `json:"date"`
	Time                  string `json:"time"`
	OperationAmount       int64  `json:"operation_amount"`
	AccountNumber         int64  `json:"account_number"`
	ReceiverAccountNumber int64  `json:"receiver_account_number"`
	//AvailableLimit        int64  `json:"available_limit"`
}

func TransactionJson(db *sql.DB) {
	transactionsJSON := []Transaction{}
	doc, err := ioutil.ReadFile("JSON/transactionjs.json")
	if err != nil {
		log.Fatal(err)

		return
	}
	err = json.Unmarshal(doc, &transactionsJSON)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(transactionsJSON)
}
