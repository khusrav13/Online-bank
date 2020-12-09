package db

const AddATM = `insert into atm(address) values(($1))`

const AddTransaction = `insert into transactionMoney(date, time, operationAmount, accountNumber, receiverAccountNumber, availableLimit) 
values(($1),($2),($3),($4),($5),($6))`

const AddUser = `insert into user(name, surname, age, gender, admin, login, password) 
values(($1),($2),($3),($4),($5),($6),($7))`

const SelectUser = `select *from users where login = ($1) and password = ($2)`

const SelectAccount = `select * from account where account.userId = ($1)`

const SelectAmount = `select amount from account where number = ($1)`

const SelectAccountNumber = `select number from account where number = ($1)`

const UpdateAccountAmountOfGiver = `update account set amount = amount - ($1) where number = ($2)`

const UpdateAccountAmountOfGainer = `update account set amount = amount + ($1) where number = ($2)`

const SelectATM = `select * from atm`

const SelectTransactionHistory = `select * from transactionMoney where accountNumber = ($1)`

const AddNewAccount = `insert into account(userId, name, number, amount, currency)
values(($1),($2),($3),($4),($6))`
