package models

import (
	"database/sql"
	"html/template"
	"time"

	"github.com/russross/blackfriday"
)

type Page struct {
	Id          int
	Title       string
	Content     string
	HTMLContent template.HTML
	UpdatedAt   int64
}

func (p Page) LastUpdate() string {
	t := time.Unix(p.UpdatedAt, 0)
	return t.Format(time.RFC1123Z)
}

func FindPage(db *sql.DB, id int) (Page, error) {
	var title []byte
	var input string
	var updated_at int64

	row := db.QueryRow("SELECT title, content, updated_at FROM pages WHERE id=?", id)
	err := row.Scan(&title, &input, &updated_at)
	if err != nil {
		return Page{}, err
	}
	output := blackfriday.MarkdownCommon([]byte(input))
	return Page{Id: id,Title: string(title),Content: input,HTMLContent: template.HTML(output),UpdatedAt: updated_at}, nil
}

func AllPages(db *sql.DB) ([]Page, error) {
	var pages []Page
	var id int
	var title []byte
	var input []byte
	var updated_at int64

	rows, err := db.Query("SELECT id, title, content, updated_at FROM pages ORDER BY updated_at DESC")
	if err != nil {
		return pages, err
	}
	for rows.Next() {
		err = rows.Scan(&id, &title, &input, &updated_at)
		if err != nil {
			return pages, err
		}
		output := blackfriday.MarkdownCommon(input)
		pages = append(pages, Page{Id: id, Title: string(title), HTMLContent: template.HTML(output), UpdatedAt: updated_at})
	}
	return pages, nil
}
