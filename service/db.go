package service

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func db(path string) *sql.DB {
	DataBase, errConnect := sql.Open("postgres", path)
	if errConnect != nil {
		log.Fatalln(errConnect)
	}
	errPing := DataBase.Ping()
	if errPing != nil {
		log.Fatalln(errPing)
	}
	return DataBase
}
