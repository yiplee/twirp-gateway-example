package book

import (
	"context"

	"github.com/patrickmn/go-cache"
	"github.com/twitchtv/twirp"
	"github.com/yiplee/twirp-gateway-example/core"
)

func New() core.BookStore {
	return &bookStore{c: cache.New(cache.NoExpiration, cache.NoExpiration)}
}

type bookStore struct {
	c *cache.Cache
}

func (b *bookStore) Create(ctx context.Context, book *core.Book) error {
	b.c.SetDefault(book.ID, book)
	return nil
}

func (b *bookStore) Find(ctx context.Context, id string) (*core.Book, error) {
	if item, ok := b.c.Get(id); ok {
		return item.(*core.Book), nil
	}

	return nil, twirp.NotFoundError("book not found")
}
