package services

import (
	"SecondProject/db"
	"SecondProject/models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const Authorization = `1.Sign in
2.Exit`

func NewAuthorization(Db *sql.DB) (user models.User) {
	var login, password string
	fmt.Println("Login:")
	_, err := fmt.Scan(&login)
	if err != nil {
		log.Println("Incorrect password", err)
	}
	fmt.Println("Password:")
	_, err = fmt.Scan(&password)
	if err != nil {
		log.Println("Incorrect password", err)
	}
	row := Db.QueryRow(db.SelectUser, login, password)
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Age,
		&user.Gender,
		&user.Admin,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		log.Println("You have not sign up!!")
		os.Exit(1)
	}
	return
}

func Authorizationss(Db *sql.DB) {
	fmt.Println(Authorization)
	var value int
	_, err := fmt.Scan(&value)
	if err != nil {
		log.Println("Try again", err)
	}
	switch value {
	case 1:
		User := NewAuthorization(Db)
		if User.Admin == true {
			Admins(Db, User)
		} else {
			Users(Db, User)
		}
	case 2:
		os.Exit(1)
	default:
		log.Fatal("Incorrect format please try again")
	}
}
