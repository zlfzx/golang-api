package entities

type Book struct {
	ID          int    `db:"id"`
	AuthorID    int    `db:"author_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Year        int    `db:"year"`
	Pages       int    `db:"pages"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
