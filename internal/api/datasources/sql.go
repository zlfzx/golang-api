package datasources

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Prepare(db *sqlx.DB, query string) *sqlx.Stmt {
	s, err := db.Preparex(query)
	if err != nil {
		// log error stmt
		fmt.Println("error preparing statement")
	}

	return s
}

func PrepareNamed(db *sqlx.DB, query string) *sqlx.NamedStmt {
	s, err := db.PrepareNamed(query)
	if err != nil {
		// log error stmt
		fmt.Println("error preparing statement")
	}

	return s
}
