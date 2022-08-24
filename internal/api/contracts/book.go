package contracts

import "golang-api/internal/api/entities"

type BookRepository interface {
	Find() (books []entities.Book, err error)
	FindById(id int) (book entities.Book, err error)
	Insert(book entities.Book) (err error)
	Update(book entities.Book) (err error)
	Delete(id int) (err error)
}

type BookService interface {
	Get() (books []entities.Book, err error)
	GetById(id int) (book entities.Book, err error)
	Insert(book entities.Book) (err error)
	Update(book entities.Book) (err error)
	Delete(id int) (err error)
}
