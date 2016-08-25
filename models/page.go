package models

import (
	"database/sql"
	"html/template"
	"strconv"
	"time"

	"github.com/russross/blackfriday"
)

type Page struct {
	ID          int
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
	if p.ID != 0 {
		sql := "UPDATE pages SET title=?, content=?, updated_at=? WHERE id=?"
		_, err = db.Exec(sql, p.Title, p.Content, now, p.ID)
	} else {
		sql := "INSERT INTO pages (title,content, updated_at) VALUES (?, ?, ?)"
		_, err = db.Exec(sql, p.Title, p.Content, now)
	}
	return err
}

func FindPage(db *sql.DB, id int) (Page, error) {
	var title []byte
	var input string
	var updatedAt int64

	sql := "SELECT title, content, updated_at FROM pages WHERE id=?"
	row := db.QueryRow(sql, id)
	err := row.Scan(&title, &input, &updatedAt)
	if err != nil {
		return Page{}, err
	}
	output := blackfriday.MarkdownCommon([]byte(input))
	page := Page{
		ID:          id,
		Title:       string(title),
		Content:     input,
		HTMLContent: template.HTML(output),
		UpdatedAt:   updatedAt,
	}
	return page, nil
}

func AllPages(db *sql.DB) ([]Page, error) {
	var pages []Page
	var id int
	var title []byte
	var input []byte
	var updatedAt int64

	sql := "SELECT id, title, content, updated_at FROM pages ORDER BY updated_at DESC"
	rows, err := db.Query(sql)
	if err != nil {
		return pages, err
	}
	for rows.Next() {
		err = rows.Scan(&id, &title, &input, &updatedAt)
		if err != nil {
			return pages, err
		}
		output := blackfriday.MarkdownCommon(input)
		page := Page{
			ID:          id,
			Title:       string(title),
			HTMLContent: template.HTML(output),
			UpdatedAt:   updatedAt,
		}
		pages = append(pages, page)
	}
	return pages, nil
}

func Search(db *sql.DB, query string) ([]Page, error) {
	var pages []Page
	var id int
	var title []byte
	var input []byte
	var updatedAt int64

	sql := `
		SELECT
			id, title, content, updated_at
		FROM
			pages
		WHERE
			title like ? or content like ?
		ORDER BY updated_at DESC`
	rows, err := db.Query(sql, "%"+query+"%", "%"+query+"%")
	if err != nil {
		return pages, err
	}
	for rows.Next() {
		err = rows.Scan(&id, &title, &input, &updatedAt)
		if err != nil {
			return pages, err
		}
		output := blackfriday.MarkdownCommon(input)
		page := Page{
			ID:          id,
			Title:       string(title),
			HTMLContent: template.HTML(output),
			UpdatedAt:   updatedAt,
		}
		pages = append(pages, page)
	}
	return pages, nil
}
