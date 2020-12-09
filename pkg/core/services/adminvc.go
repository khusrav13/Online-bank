package services

import (
	"SecondProject/db"
	"SecondProject/models"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

const Admin = `1. Address ATM
2.Add new ATM
3.Show Balance
4.Payment
5.Transaction History
6.List of Accounts
7.List of ATMs in the city
8.List of transactions
9.List of users
10.Exit`

func Admins(Db *sql.DB, user models.User) {
	fmt.Println(Admin)
	var value int
	_, err := fmt.Scan(&value)
	if err != nil {
		log.Println("Please try again", err)
	}
	switch value {
	case 1:
		ATMsAddresses(Db, user) //1 )
	case 2:
		AddNewATM(Db, user) //2
	case 3:
		ShowAmount(Db, user) ///3
	case 4:
		Payment(Db, user) ///4
	case 5:
		ShowHistoryOfTransaction(Db, user)
	case 6:
		models.AccountJson(Db)
	case 7:
		models.AtmJson(Db)
	case 8:
		models.TransactionJson(Db)
	case 9:
		models.UsersJson(Db)
	case 10:
		os.Exit(1)
	default:
		fmt.Println("Please try again")
	}
}

func AddNewATMS(Db *sql.DB, address string) (err error) {
	_, err = Db.Exec(db.AddATM, address)
	if err != nil {
		return err
	}
	return
}

func AddNewATM(Db *sql.DB, user models.User) {
	var h string
	_, err := fmt.Scan(&h)
	if err != nil {
		log.Println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	address, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Can't read command:", err)
	}
	fmt.Println(h)
	sprint := fmt.Sprintf("%s %s", h, address)
	fmt.Println(sprint)
	err = AddNewATMS(Db, sprint)
	if err != nil {
		return
	}
	fmt.Println("Successfully Added")
}
