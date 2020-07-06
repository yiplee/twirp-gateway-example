package book

import (
	context "context"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
	"github.com/gofrs/uuid"
	"github.com/twitchtv/twirp"
	"github.com/yiplee/twirp-gateway-example/core"
	"github.com/yiplee/twirp-gateway-example/handler/codes"
	"github.com/yiplee/twirp-gateway-example/handler/rpc/views"
	"github.com/yiplee/twirp-gateway-example/proto"
)

func New(books core.BookStore) proto.TwirpServer {
	s := &service{books: books}
	return proto.NewBookServiceServer(s, nil)
}

type service struct {
	books core.BookStore
}

func (s *service) Create(ctx context.Context, create *proto.BookServiceReq_Create) (*proto.Book, error) {
	if utf8.RuneCountInString(create.Name) < 6 {
		err := twirp.InvalidArgumentError("name", "too short")
		return nil, codes.With(err, codes.BookNameTooShort)
	}

	if !govalidator.IsURL(create.Cover) {
		return nil, twirp.InvalidArgumentError("cover", "must be valid url")
	}

	book := &core.Book{
		ID:    uuid.Must(uuid.NewV4()).String(),
		Name:  create.Name,
		Cover: create.Cover,
	}

	if err := s.books.Create(ctx, book); err != nil {
		return nil, err
	}

	return views.Book(book), nil
}

func (s *service) Find(ctx context.Context, find *proto.BookServiceReq_Find) (*proto.Book, error) {
	book, err := s.books.Find(ctx, find.Id)
	if err != nil {
		return nil, err
	}

	return views.Book(book), nil
}
