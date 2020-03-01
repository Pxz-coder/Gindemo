package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

)

//数据库常量
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pan123"
	dbname   = "GinDemo"
)


//连接数据库
func Connect()(*sql.DB){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
