package models

import (
	"database/sql"
	"html/template"

	"github.com/russross/blackfriday"
)

type Page struct {
	Id          int
	Title       string
	Content     string
	HTMLContent template.HTML
}

func FindPage(db *sql.DB, id int) (Page, error) {
	var title []byte
	var input string

	row := db.QueryRow("SELECT title, content FROM pages WHERE id=?", id)
	err := row.Scan(&title, &input)
	if err != nil {
		return Page{}, err
	}
	output := blackfriday.MarkdownCommon([]byte(input))
	return Page{Id: id, Title: string(title), Content: input, HTMLContent: template.HTML(output)}, nil
}

func AllPages(db *sql.DB) ([]Page, error) {
	var pages []Page
	var id int
	var title []byte
	var input []byte

	rows, err := db.Query("SELECT * FROM pages")
	if err != nil {
		return pages, err
	}
	for rows.Next() {
		err = rows.Scan(&id, &title, &input)
		if err != nil {
			return pages, err
		}
		output := blackfriday.MarkdownCommon(input)
		pages = append(pages, Page{Id: id, Title: string(title), HTMLContent: template.HTML(output)})
	}
	return pages, nil
}
