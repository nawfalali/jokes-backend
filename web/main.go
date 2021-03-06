package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"nawfalali.com/jokesapp/pkg/models/mysql"
	//	"net/http"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	jokes    *mysql.JokeModel
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "nfl:Asdf!234@/jokes?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)

	if err != nil {
		errorLog.Fatal(err)
	} else {
		infoLog.Printf("DB connection established")
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		jokes:    &mysql.JokeModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Call the new app.routes() method
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

	fmt.Println(addr)

}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}
