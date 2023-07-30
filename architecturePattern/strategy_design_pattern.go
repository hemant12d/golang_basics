package main

import "fmt"

// The Strategy pattern basically allows an object to select and execute at runtime without knowing how to implement
// the business logic it wants.

type IDBConnecton interface {
	connect()
}

// Created the SqlConnection struct, and for implementing to IDBconnect interface
type sqlConnection struct {
	connectionUrl string
}

func (sqlC sqlConnection) connect() {
	fmt.Println("Connecting with SQL database...")
}

type mongoConnection struct {
	connectionUrl string
}

func (sqlC mongoConnection) connect() {
	fmt.Println("Connecting with SQL database...")
}

type oracleConnection struct {
	connectionUrl string
}

func (sqlC oracleConnection) connect() {
	fmt.Println("Connecting with SQL database...")
}

type DBConnection struct {
	db IDBConnecton
}

func main() {
	sqlConn := sqlConnection{connectionUrl: "http://localhost:3306"}
	con1 := DBConnection{db: sqlConn}
	con1.db.connect()

	mongoConn := mongoConnection{connectionUrl: "http://localhost:27017"}
	con2 := DBConnection{db: mongoConn}
	con2.db.connect()
}
