package models

import (
	"database/sql"
	"html/template"
	"strconv"
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
	return t.Format("02-01-2006 15:04")
}

func (p *Page) Save(db *sql.DB) error {
	var err error
	now := strconv.FormatInt(time.Now().Unix(), 10)
	if p.Id != 0 {
		_, err = db.Exec("UPDATE pages SET title=?, content=?, updated_at=? WHERE id=?", p.Title, p.Content, now, p.Id)
	} else {
		_, err = db.Exec("INSERT INTO pages (title,content, updated_at) VALUES (?, ?, ?)", p.Title, p.Content, now)
	}
	return err
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
	return Page{Id: id, Title: string(title), Content: input, HTMLContent: template.HTML(output), UpdatedAt: updated_at}, nil
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

func Search(db *sql.DB, query string) ([]Page, error) {
	var pages []Page
	var id int
	var title []byte
	var input []byte
	var updated_at int64

	rows, err := db.Query("SELECT id, title, content, updated_at FROM pages WHERE title like ? or content like ? ORDER BY updated_at DESC", "%"+query+"%", "%"+query+"%")
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
