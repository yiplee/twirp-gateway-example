package views

import (
	"github.com/yiplee/twirp-gateway-example/core"
	"github.com/yiplee/twirp-gateway-example/proto"
)

func Book(book *core.Book) *proto.Book {
	return &proto.Book{
		Id:    book.ID,
		Name:  book.Name,
		Cover: book.Cover,
	}
}
