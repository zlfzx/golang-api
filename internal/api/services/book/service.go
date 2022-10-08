package book

import (
	"golang-api/internal/api/contracts"
	"golang-api/internal/api/entities"
)

type Service struct {
	repo contracts.BookRepository
}

func Init(app *contracts.App) (svc contracts.BookService) {

	r := initRepository(app.Datasource.WriterDB, app.Datasource.ReaderDB)

	svc = &Service{
		repo: r,
	}

	return
}

func (s Service) Get() (books []entities.Book, err error) {
	books, err = s.repo.Find()
	return
}

func (s Service) GetById(id int) (book entities.Book, err error) {
	book, err = s.repo.FindById(id)
	return
}

func (s Service) Insert(book entities.Book) (err error) {
	err = s.repo.Insert(book)
	return
}

func (s Service) Update(book entities.Book) (err error) {
	err = s.repo.Update(book)
	return
}

func (s Service) Delete(id int) (err error) {
	err = s.repo.Delete(id)
	return
}
