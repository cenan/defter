package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cenan/defter/controllers"
	"github.com/skratchdot/open-golang/open"
	"gopkg.in/ini.v1"
)

func setupRoutes(db *sql.DB) {
	http.Handle("/", controllers.IndexPage(db))
	http.Handle("/notebook", controllers.NotebookPage(db))
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg, err := ini.Load("config.ini")
	checkError(err)
	verboseOutput, err := cfg.Section("general").Key("verbose").Bool()
	checkError(err)
	dbPath := cfg.Section("db").Key("path").String()
	port, err := cfg.Section("server").Key("port").Int()
	checkError(err)
	if verboseOutput == false {
		log.SetOutput(ioutil.Discard)
	}
	db, err := sql.Open("sqlite3", dbPath)
	checkError(err)
	defer db.Close()
	startWebClient(db, port)
}
