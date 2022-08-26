package database

import "database/sql"

const Default_Host = "localhost"
const Default_Port = 5432
const Default_User = "admin"
const Default_Password = "password"
const Default_Dbname = "postgres"

type Database struct {
	Instance *sql.DB
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func NewDatabase(host string, port int, user string, password string, dbname string) *Database {
	return &Database{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

type DatabaseID int64
