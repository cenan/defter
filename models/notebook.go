package models

import "database/sql"

type Notebook struct {
	ID   int
	Name string
}

func AllNotebooks(db *sql.DB) ([]Notebook, error) {
	var notebooks []Notebook
	var id int
	var name []byte

	sql := "SELECT id, name FROM notebooks ORDER BY id"
	rows, err := db.Query(sql)
	if err != nil {
		return notebooks, err
	}
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return notebooks, err
		}
		notebook := Notebook{
			ID:   id,
			Name: string(name),
		}
		notebooks = append(notebooks, notebook)
	}
	return notebooks, nil
}
