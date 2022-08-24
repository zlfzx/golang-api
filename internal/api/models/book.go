package models

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
