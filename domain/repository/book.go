package repository

import "github.com/yyh-gl/go-api-server-by-ddd/domain/model"

type Book interface {
	GetAll() *[]model.Book
}
