package contracts

import "golang-api/internal/api/entities"

type AuthorRepository interface {
	Find() (authors []entities.Author, err error)
	FindById(id int) (author entities.Author, err error)
	Insert(author entities.Author) (err error)
	Update(author entities.Author) (err error)
	Delete(id int) (err error)
}

type AuthorService interface {
	Get() (authors []entities.Author, err error)
	GetById(id int) (author entities.Author, err error)
	Insert(author entities.Author) (err error)
	Update(author entities.Author) (err error)
	Delete(id int) (err error)
}
