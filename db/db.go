package db

import (
	"log"

	r"gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/mkn2016/rss-feed-parser/structs"
)

//GetDbConnection ... RethinkDB Database Connection
func GetDbConnection() *r.Session{
	var err error
	
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		MaxOpen: 10,
		InitialCap: 10,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Connection established successfully")
	return session
}

//ListDatabases ... Lists the databases in rethinkdb
func ListDatabases() []interface{}{
	var databases []interface{}
	log.Println("Getting all databases in rethinkdb cluster.")
	session := GetDbConnection()

	res, err := r.DBList().Run(session)
	defer res.Close()

	if err != nil {
		log.Println("Could not list databases in rethinkdb cluster.")
	}

	err = res.All(&databases)
	if err != nil {
		log.Println("error occured")
	}
	log.Println("Successfully got a list of all databases in rethinkdb cluster.", databases)
	return databases
}

//ListTables ... Lists the tables in rethinkdb
func ListTables(database string) []interface{}{
	var tables []interface{}
	log.Println("Getting all databases in rethinkdb cluster.")
	session := GetDbConnection()

	res, err := r.DB(database).TableList().Run(session)
	defer res.Close()

	if err != nil {
		log.Println("Could not list databases in rethinkdb cluster.")
	}

	err = res.All(&tables)
	if err != nil {
		log.Println("Error occured")
	}
	log.Println("Successfully got a list of all databases in rethinkdb cluster.", tables)
	return tables
}


//CreateDatabase ... creates a database in rethinkdb
func CreateDatabase(database string) {
	log.Println("Creating Database: ", database, "in rethinkdb cluster.")
	session := GetDbConnection()

	databases := ListDatabases()

	for _, v := range databases{
		if v == database {
			log.Println("Skipping database creation. Database already exists")
		}
	}
	res, err := r.DBCreate(database).Run(session)
	defer res.Close()

	if res.Err() != nil {
		log.Println("Could not create database: ", database, "in rethinkdb cluster.", err)
	}
	log.Println("Successfully created Database: ", database, "in rethinkdb cluster.")
}

//CreateTable ... creates a table in rethinkdb
func CreateTable(database string, table string) {
	log.Println("Creating table: ", table, "in rethinkdb cluster.")
	session := GetDbConnection()

	tables := ListTables(database)

	for _, v := range tables{
		if v == table {
			log.Println("Skipping table creation. Table already exists")
		}
	}

	res, err := r.DB(database).TableCreate(table).Run(session)
	defer res.Close()

	if err != nil {
		log.Println("Could not create table: ", table, "in rethinkdb cluster.")
	}
	log.Println("Successfully created table: ", table, "in rethinkdb cluster.")
}

//InsertRecordToTable ... insert a document into table
func InsertRecordToTable(database string, table string, document []structs.RssParent) {
	log.Println("Inserting document to table: ", table, "in database: ", database)
	session := GetDbConnection()
	err := r.DB(database).Table(table).Insert(document).Exec(session)
	if err != nil {
		log.Println("Error occured while executing insert to table: ", table, "in database: ", database)
		log.Println("Error is: ", err)
	}
	log.Println("Successfully inserted document to table: ", table, "in database: ", database)
}

// GetTableRecords ... gets all the records stored in rethinkdb table
func GetTableRecords(database string, table string) ([]structs.RssParent){
	result := []structs.RssParent{}
	session := GetDbConnection()
	
	res, err := r.DB(database).Table(table).OrderBy("id").Run(session)

	if err != nil {
		log.Println("Nothing to show because i am empty")
		return []structs.RssParent{}
	}

	err = res.All(&result)
	return result
}

//GetRecordByID ... get a speficic record from table by ID
func GetRecordByID(database string, table string, id string) structs.RssParent{
	result := structs.RssParent{}

	session := GetDbConnection()

	res, err := r.DB(database).Table(table).Get(id).Run(session)

	if err != nil {
		log.Println("Error occured: ", err)
	}
	res.One(&result)
	return result
}
