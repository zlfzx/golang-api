package author

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

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) contracts.AuthorRepository {
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

func (r Repository) Find() (authors []entities.Author, err error) {
	err = r.stmt.findList.Select(&authors)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("fail to get list , err: %v", err))
	}
	return
}

func (r Repository) FindById(id int) (author entities.Author, err error) {
	row := r.stmt.findById.QueryRow(id)
	err = row.Scan(
		&author.ID,
		&author.Name,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if err != nil {
		alog.Logger.Error(fmt.Errorf("fail to get guest by id, err: %v", err))
	}
	return
}

func (r Repository) Insert(author entities.Author) (err error) {
	_, err = r.stmt.insert.Exec(author)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("error while inserting guest, err: %v", err))
	}

	return err
}

func (r Repository) Update(author entities.Author) (err error) {
	_, err = r.stmt.update.Exec(author)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("error while updating guest, err: %v", err))
	}

	return err
}

func (r Repository) Delete(id int) (err error) {
	_, err = r.stmt.delete.Exec(id)
	if err != nil {
		alog.Logger.Error(fmt.Errorf("error while deleting guest, err: %v", err))
	}

	return err
}
