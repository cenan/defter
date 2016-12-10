package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cenan/defter/models"
)

func IndexPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pages, err := models.AllPages(db)
		notebooks, err := models.AllNotebooks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t := getTemplate("index")
		err = t.ExecuteTemplate(w, "base", notebooks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func NotebookPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pages, err := models.FindPagesOfNotebook(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t := getTemplate("notebook")
		err = t.ExecuteTemplate(w, "base", pages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func SearchPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pages, err := models.Search(db, r.URL.Query().Get("query"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t := getTemplate("index")
		err = t.ExecuteTemplate(w, "base", pages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func NewPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := getTemplate("new")
		err := t.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func CreatePage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notebookID, err := strconv.Atoi(r.FormValue("notebook_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		page := models.Page{
			Title:      r.FormValue("title"),
			Content:    r.FormValue("content"),
			NotebookID: notebookID,
		}
		err = page.Save(db)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func ShowPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p, err := models.FindPage(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t := getTemplate("show")
		err = t.ExecuteTemplate(w, "base", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func EditPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p, err := models.FindPage(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t := getTemplate("edit")
		err = t.ExecuteTemplate(w, "base", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func SavePage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		page := models.Page{
			ID:      id,
			Title:   r.FormValue("title"),
			Content: r.FormValue("content"),
		}
		err = page.Save(db)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, fmt.Sprintf("/show?id=%d", id), http.StatusFound)
	}
}

func Close() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	}
}

func getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles("templates/base.html", "templates/"+templateName+".html")
	return template.Must(t, err)
}
