package db

const CreateUser = `CREATE TABLE IF NOT EXISTS users (
	Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	name TEXT NOT NULL,
	surname TEXT NOT NULL,
	age INTEGER NOT NULL,
	gender TEXT NOT NULL,
	admin boolean not null,
	login TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	remove BOOLEAN NOT NULL DEFAULT FALSE
)`

const CreateATM = `CREATE TABLE IF NOT EXISTS atm (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	address TEXT NOT NULL,
	status BOOLEAN NOT NULL DEFAULT TRUE
)`

const CreateTransaction = `CREATE TABLE IF NOT EXISTS transactionMoney (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	date TEXT NOT NULL,
	time TEXT NOT NULL,
	operationAmount INTEGER NOT NULL,
	accountNumber INTEGER NOT NULL,
	receiverAccountNumber INTEGER NOT NULL,
	availableLimit INTEGER NOT NULL
)`

const CreateAccount = `create table if not exists account (
	Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	userId INTEGER REFERENCES user(Id),
	name TEXT NOT NULL,
	number TEXT NOT NULL UNIQUE,
	amount INTEGER NOT NULL,
	currency TEXT NOT NULL
)`
