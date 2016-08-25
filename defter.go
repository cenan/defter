package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cenan/defter/controllers"
	"github.com/skratchdot/open-golang/open"
)

func setupRoutes(db *sql.DB) {
	http.Handle("/", controllers.IndexPage(db))
	http.Handle("/search", controllers.SearchPage(db))
	http.Handle("/new", controllers.NewPage(db))
	http.Handle("/create", controllers.CreatePage(db))
	http.Handle("/show", controllers.ShowPage(db))
	http.Handle("/edit", controllers.EditPage(db))
	http.Handle("/save", controllers.SavePage(db))
	http.Handle("/close", controllers.Close())

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func startWebClient(db *sql.DB, port int) {
	setupRoutes(db)
	log.Printf("Started serving on port %d", port)
	open.Run(fmt.Sprintf("http://localhost:%d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	verboseOutput := flag.Bool("verbose", true, "verbose output")
	port := flag.Int("port", 5000, "server port")
	dbPath := flag.String("db", "/Users/cenan/Dropbox/defter.sqlite", "database file")
	flag.Parse()

	if *verboseOutput == false {
		log.SetOutput(ioutil.Discard)
	}

	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		panic("Cannot open database: " + *dbPath)
	}
	defer db.Close()
	startWebClient(db, *port)
}
