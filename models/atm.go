package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ATM struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
	Status  bool   `json:"status"`
}

//////////Function for JSON
func AtmJson(db *sql.DB) {
	atmsJSON := []ATM{}
	doc, err := ioutil.ReadFile("JSON/atmjs.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(doc, &atmsJSON)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(atmsJSON)
}
