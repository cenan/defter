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
)

func startWebClient(db *sql.DB, port int) {
	http.Handle("/", controllers.IndexPage(db))
	http.Handle("/new", controllers.NewPage(db))
	http.Handle("/create", controllers.CreatePage(db))
	http.Handle("/show", controllers.ShowPage(db))
	http.Handle("/edit", controllers.EditPage(db))
	http.Handle("/save", controllers.SavePage(db))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Printf("Started serving on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	verboseOutput := flag.Bool("verbose", false, "verbose output")
	port := flag.Int("port", 5000, "server port")
	flag.Parse()

	if *verboseOutput == false {
		log.SetOutput(ioutil.Discard)
	}

	const db_path = "./db/defter.sqlite"
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic("Cannot open database: " + db_path)
	}
	defer db.Close()
	startWebClient(db, *port)
}
