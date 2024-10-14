package repository

import (
	"database/sql"
	"log"
	_"github.com/lib/pq"
)
var DB *sql.DB

func InitDB() {
	var err error
	dsn := "user=jaanar-blrm19 password=Anti@19icej dbname=blog_db sslmode=disable"


	DB,err = sql.Open("postgres",dsn)
	if err !=nil{
		log.Fatalf("Failed to connect to the  database: %v",err)
	}

	if err = DB.Ping(); err!= nil{
		log.Fatalf("Failed to ping the database: %v",err)
	}
	log.Println("Connected to the database")
}

