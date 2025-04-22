package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR,
	comment TEXT,
	repeat VARCHAR
);

CREATE INDEX idx_scheduler_date ON scheduler(date);
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	install := os.IsNotExist(err)

	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}

	if install {
		_, err = DB.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}
