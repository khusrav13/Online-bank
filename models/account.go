package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Account struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Name     string `json:"name"`
	Number   int64  `json:"number"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func AccountJson(db *sql.DB) {
	accountJSON := []Account{}
	doc, err := ioutil.ReadFile("JSON/accountjs.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(doc, &accountJSON)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(accountJSON)
}
