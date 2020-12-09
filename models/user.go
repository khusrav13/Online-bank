package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
	Admin    bool   `json:"admin"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Remove   bool   `json:"remove"`
}

//////////Function for JSON
func UsersJson(db *sql.DB) {
	userJSON := []User{}
	doc, err := ioutil.ReadFile("JSON/userjs.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(doc, &userJSON)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(userJSON)
}
