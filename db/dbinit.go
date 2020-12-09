package db

import (
	"database/sql"
	"fmt"
)

func Dbinits(db *sql.DB) (err error) {
	var i int64 = 1
	DDLs := []string{CreateATM, CreateUser, CreateTransaction, CreateAccount}
	for _, ddl := range DDLs {
		_, err = db.Exec(ddl)
		if err != nil {
			fmt.Println(i, err)
		}
		i++
	}
	return
}
