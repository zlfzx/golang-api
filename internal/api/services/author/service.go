package author

import (
	"golang-api/internal/api/contracts"
	"golang-api/internal/api/entities"
)

type Service struct {
	repo contracts.AuthorRepository
}

func Init(app *contracts.App) (svc contracts.AuthorService) {

	r := initRepository(app.Datasource.WriterDB, app.Datasource.ReaderDB)

	svc = &Service{
		repo: r,
	}

	return
}

func (s Service) Get() (authors []entities.Author, err error) {
	authors, err = s.repo.Find()
	return
}

func (s Service) GetById(id int) (author entities.Author, err error) {
	author, err = s.repo.FindById(id)
	return
}

func (s Service) Insert(author entities.Author) (err error) {
	err = s.repo.Insert(author)
	return
}

func (s Service) Update(author entities.Author) (err error) {
	err = s.repo.Update(author)
	return
}

func (s Service) Delete(id int) (err error) {
	err = s.repo.Delete(id)
	return
}
