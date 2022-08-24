package models

import (
	"errors"
)

type Author struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CallbackResponse struct {
	CallbackResponse Author
}

func (author Author) Validate() error {
	if author.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
