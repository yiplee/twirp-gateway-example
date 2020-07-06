package core

import (
	"context"
)

type Book struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Cover string `json:"cover,omitempty"`
}

type BookStore interface {
	Create(ctx context.Context, book *Book) error
	Find(ctx context.Context, id string) (*Book, error)
}
