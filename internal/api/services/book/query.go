package book

const (
	insert = `
		INSERT INTO books (
			name
		) VALUES (
			:name
		)
	`

	update = `
		UPDATE books
		SET 
			name = :name
		WHERE
			id = :id
	`

	findById = `SELECT * FROM books WHERE id = ?`

	findList = `SELECT * FROM books ORDER BY title ASC`

	delete = `DELETE FROM books WHERE id = ?`
)
