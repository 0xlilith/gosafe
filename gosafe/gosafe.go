package gosafe

/*
 * gosafe database API
 * author: 0xlilith
 */

import "database/sql"

type Locker struct {
	DB *sql.DB
}

func (locker *Locker) Get() []Item {
	items := []Item{}
	rows, err := locker.DB.Query(`
		SELECT * from gosafe 
	`)
	if err != nil {
		panic("failed to connect sql server: " + err.Error())
	}
	var id int
	var name string
	var password string
	for rows.Next() {
		rows.Scan(&id, &name, &password)
		item := Item{
			ID:       id,
			Name:     name,
			Password: password,
		}
		items = append(items, item)
	}
	return items
}

func (locker *Locker) Add(item Item) {
	stmt, err := locker.DB.Prepare(`
		INSERT INTO gosafe (item, password) values (?, ?) 
	`)
	if err != nil {
		panic("failed to connect sql server: " + err.Error())
	}
	stmt.Exec(item.Name, item.Password)
}

func NewPass(db *sql.DB) *Locker {

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "gosafe" (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"item"	TEXT,
			"password"	TEXT
		);
	`)
	if err != nil {
		panic("ping error: " + err.Error())
	}
	stmt.Exec()

	return &Locker{
		DB: db,
	}
}
