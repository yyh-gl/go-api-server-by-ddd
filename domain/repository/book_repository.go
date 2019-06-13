package repository

import "github.com/yyh-gl/go-api-server-by-ddd/domain/model"

type BookRepository interface {
	GetAll() []*model.Book
}
