package author

const (
	insert = `
		INSERT INTO authors (
			name
		) VALUES (
			:name
		)
	`

	update = `
		UPDATE authors
		SET 
			name = :name
		WHERE
			id = :id
	`

	findById = `SELECT * FROM authors WHERE id = ?`

	findList = `SELECT * FROM authors ORDER BY name ASC`

	delete = `DELETE FROM authors WHERE id = ?`
)
