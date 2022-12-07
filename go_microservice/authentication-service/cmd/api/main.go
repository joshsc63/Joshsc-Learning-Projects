package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Can be same port as backend. Docker allows same port across mulitple containers
const webPort = "80"

var counts int64

type Config struct {
	DB *sql.DB
	Models data.Models
}

// web API service 
func main() {
	log.Println("starting authentication service")
	
	// connect to DB 
	conn := connectToDB()
	if conn == nil {
		log.Panic("Cant connect to Postgres DB ")
	}
	
	app := Config{
		DB: conn,
		Models: data.New(conn),
	}
	
	//Setup webserver
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	
}

//dsn = connection string from env
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

// DB may be up b4 service. Have service connect
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++ 
		} else {
			// SUCCESS
			log.Panicln("Connected to Postgres DB")
			return connection
		}
		
		if counts > 10 {
			log.Println(err)
			return nil
		}
		
		log.Println("Pausing for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}