package book

import (
	"fmt"
	"golang-api/internal/api/contracts"
	"golang-api/internal/api/datasources"
	"golang-api/internal/api/entities"
	"golang-api/pkg/alog"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	insert   *sqlx.NamedStmt
	update   *sqlx.NamedStmt
	findById *sqlx.Stmt
	findList *sqlx.Stmt
	delete   *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) contracts.BookRepository {
	stmts := Statement{
		insert:   datasources.PrepareNamed(dbWriter, insert),
		update:   datasources.PrepareNamed(dbWriter, update),
		findById: datasources.Prepare(dbReader, findById),
		findList: datasources.Prepare(dbReader, findList),
		delete:   datasources.Prepare(dbWriter, delete),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) Find() (books []entities.Book, err error) {
	err = r.stmt.findList.Select(&books)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("fail to get list , err: %v", err))
	}
	return
}

func (r Repository) FindById(id int) (book entities.Book, err error) {
	row := r.stmt.findById.QueryRow(id)
	err = row.Scan(
		&book.ID,
		&book.Title,
		&book.AuthorID,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		alog.Logger.Error(fmt.Errorf("fail to get book by id, err: %v", err))
	}
	return
}

func (r Repository) Insert(book entities.Book) (err error) {
	_, err = r.stmt.insert.Exec(book)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("fail to insert book, err: %v", err))
	}

	return err
}

func (r Repository) Update(book entities.Book) (err error) {
	_, err = r.stmt.update.Exec(book)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("fail to update book, err: %v", err))
	}

	return err
}

func (r Repository) Delete(id int) (err error) {
	_, err = r.stmt.delete.Exec(id)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("fail to delete book, err: %v", err))
	}

	return err
}
