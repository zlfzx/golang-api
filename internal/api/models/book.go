package models

import "errors"

type Book struct {
	ID          int    `json:"id"`
	AuthorID    int    `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Pages       int    `json:"pages"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (book Book) Validate() error {

	if book.Title == "" {
		return errors.New("title is required")
	}

	if book.Description == "" {
		return errors.New("description is required")
	}

	if book.Year == 0 {
		return errors.New("year is required")
	}

	if book.Pages == 0 {
		return errors.New("pages is required")
	}

	return nil
}
