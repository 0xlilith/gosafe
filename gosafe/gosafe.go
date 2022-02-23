package gosafe

/*
 * gosafe database API
 * author: 0xlilith
 */

import "database/sql"

type Locker struct {
	DB *sql.DB
}

// func (locker *Locker) Get() []Item {

// }

func (locker *Locker) Add(item Item) {
	stmt, e := locker.DB.Prepare(`
		INSERT INTO gosafe (item, password) values (?, ?) 
	`)
	if e != nil {
		panic("failed to connect sql server: " + e.Error())
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
