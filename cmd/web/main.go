package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	_"github.com/go-sql-driver/mysql"
	"github.com/smg061/snippetbox/pkg/models/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *mysql.SnippetModel
}
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL database connection")
	// actuallhy parse  the flag, otherwie it will always have the value of :4000
	flag.Parse()


	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime| log.Lshortfile)
	db, err := openDB(*dsn)
	
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &mysql.SnippetModel{DB:db},
	}
	if err != nil {
		app.errorLog.Fatalln("Error connecting to database")
	}
	defer db.Close()
	
	infoLog.Printf("starting server on %s", *addr)
	srv := &http.Server{
		Handler: app.routes(),
		Addr: *addr,
		ErrorLog: errorLog,
	}
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB (dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil  {
		return nil, err
	}
	if err := db.Ping(); err !=nil {
		return nil, err
	}
	return db, nil
}