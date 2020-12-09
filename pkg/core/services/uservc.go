package services

import (
	"SecondProject/db"
	"SecondProject/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

const User = `1.Address
ATM
2.Show Balance
3.Payment
4.Transaction History
5.Exit`

func Users(DBs *sql.DB, user models.User) {
	fmt.Println(User)
	var value int
	_, err := fmt.Scan(&value)
	if err != nil {
		log.Println("Error please try again", err)
	}
	switch value {
	case 1:
		ATMsAddresses(DBs, user) //1
	case 2:
		ShowAmount(DBs, user)
	case 3:
		Payment(DBs, user)
	case 4:
		ShowHistoryOfTransaction(DBs, user)
	case 5:
		os.Exit(1)
	default:
		fmt.Println("Please try again")
	}
}

func ATMsAddresses(Db *sql.DB, user models.User) {
	rows, err := Db.Query(db.SelectATM)
	if err != nil {
		panic(err)
	}
	atms := []models.ATM{}
	for rows.Next() {
		p := models.ATM{}
		err := rows.Scan(
			&p.ID,
			&p.Address,
			&p.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		atms = append(atms, p)
	}
	for _, p := range atms {
		fmt.Println(p)
	}

}

func ShowAmount(Db *sql.DB, user models.User) {
	rows, err := Db.Query(db.SelectAccount, user.ID)
	if err != nil {
		panic(err)
	}
	accounts := []models.Account{}

	for rows.Next() {
		p := models.Account{}
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Name,
			&p.Number,
			&p.Amount,
			&p.Currency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		accounts = append(accounts, p)
	}
	if len(accounts) > 1 {
		fmt.Println("Balance on the card: ")
		var total int64
		for _, p := range accounts {
			fmt.Printf("%s include %d %s \n", p.Name, p.Amount, p.Currency)
			total += p.Amount
		}
	} else {
		fmt.Printf("Balance is %d \n", accounts[0].Amount)
	}
}
func ShowHistoryOfTransaction(Db *sql.DB, user models.User) {
	rows, err := Db.Query(db.SelectAccount, user.ID)
	if err != nil {
		panic(err)
	}
	acc := []models.Account{}
	for rows.Next() {
		p := models.Account{}
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Name,
			&p.Number,
			&p.Amount,
			&p.Currency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		acc = append(acc, p)
	}
	var t int
	for i := 0; i < len(acc); i++ {
		t = i
	}
	if t < 1 {
		rows, err := Db.Query(db.SelectTransactionHistory, acc[0].Number)
		if err != nil {
			panic(err)
		}
		arr := []models.Transaction{}
		for rows.Next() {
			p := models.Transaction{}
			err := rows.Scan(
				&p.ID,
				&p.Date,
				&p.Time,
				&p.OperationAmount,
				&p.AccountNumber,
				&p.ReceiverAccountNumber)
			if err != nil {
				fmt.Println(err)
				continue
			}
			arr = append(arr, p)
		}
		for i := 0; i < len(arr); i++ {
			fmt.Println("    Bank Of America  ", "ATM number:", arr[i].ID, "Date:", arr[i].Date, "Time:", arr[i].Time, "Sum of payment:", arr[i].OperationAmount, "Your number card:", arr[i].AccountNumber, "Receiver card:", arr[i].ReceiverAccountNumber)
			//fmt.Println("Limit:",arr[i].AvailableLimit, "\n")
		}
	} else {
		var accountNumber int64
		_, err := fmt.Scan(&accountNumber)
		if err != nil {
			panic(err)
		}
		rows, err := Db.Query(db.SelectTransactionHistory, accountNumber)
		if err != nil {
			panic(err)
		}
		arr := []models.Transaction{}
		for rows.Next() {
			p := models.Transaction{}
			err := rows.Scan(
				&p.ID,
				&p.Date,
				&p.Time,
				&p.OperationAmount,
				&p.AccountNumber,
				&p.ReceiverAccountNumber)
			if err != nil {
				fmt.Println(err)
				continue
			}
			arr = append(arr, p)
		}
		for i := 0; i < len(arr); i++ {
			fmt.Println("    Bank Of America  ", "ATM number:", arr[i].ID, "Date:", arr[i].Date, "Time:", arr[i].Time, "Sum of payment:", arr[i].OperationAmount, "Your number card:", arr[i].AccountNumber, "Receiver card:", arr[i].ReceiverAccountNumber)
			//fmt.Println("Limit:",arr[i].AvailableLimit, "\n")
		}
	}

}

func Payment(Db *sql.DB, user models.User) {
	fmt.Println("Payment operation in any card in the world")
	var AccountNumber, receiverAccountNumber, translationAmount int64
	fmt.Println("Your card numbers:")
	_, err := fmt.Scan(&AccountNumber)
	if err != nil {
		log.Print("error", err)
	}
	fmt.Println("Type sender card:")
	_, err = fmt.Scan(&receiverAccountNumber)
	if err != nil {
		log.Print("Error", err)
	}
	fmt.Println("Sum:")
	_, err = fmt.Scan(&translationAmount)
	if err != nil {
		log.Print("Error", err)
	}
	var Account, receiverAccount models.Account
	row := Db.QueryRow(db.SelectAmount, AccountNumber)
	err = row.Scan(&Account.Amount)
	if err != nil {
		panic(err)
	}
	row = Db.QueryRow(db.SelectAccountNumber, AccountNumber)
	err = row.Scan(&Account.Number)
	if err != nil {
		panic(err)
	}
	row = Db.QueryRow(db.SelectAmount, receiverAccountNumber)
	err = row.Scan(&receiverAccount.Amount)
	if err != nil {
		panic(err)
	}
	row = Db.QueryRow(db.SelectAccountNumber, receiverAccountNumber)
	err = row.Scan(&receiverAccount.Number)
	if err != nil {
		panic(err)
	}
	if Account.Amount < translationAmount {
		fmt.Println("You have not enough money!")

		return
	}
	_, err = Db.Exec(db.UpdateAccountAmountOfGiver, translationAmount, AccountNumber)
	if err != nil {
		panic(err)
	}
	_, err = Db.Exec(db.UpdateAccountAmountOfGainer, translationAmount, receiverAccountNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println("Your payment accomplished successfully!")
	err = AddTransactionHistory(Db, Account, receiverAccount, translationAmount)
	if err != nil {
		panic(err)
	}

}

func AddTransactionHistory(Db *sql.DB, myAccount, receiverAccount models.Account, operationAmount int64) (err error) {
	var check models.Transaction
	data := time.Now()
	check.Date = data.Format("02-Aug-210")
	check.Time = data.Format("22:11")
	check.OperationAmount = operationAmount
	check.AccountNumber = myAccount.Number
	//check.AvailableLimit = myAccount.Amount - operationAmount
	check.ReceiverAccountNumber = receiverAccount.Number
	_, err = Db.Exec(db.AddTransaction, check.Date, check.Time, check.OperationAmount, check.AccountNumber, check.ReceiverAccountNumber)
	if err != nil {
		panic(err)
	}
	return
}
