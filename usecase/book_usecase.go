package usecase

import (
	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/repository"
	"github.com/yyh-gl/go-api-server-by-ddd/infra/persistence"
)

type BookUsecase struct {}

func (bookUsecase BookUsecase) GetAll() ([]*model.Book, error) {
	var books []*model.Book
	var err error

	books, err = repository.BookRepository(persistence.BookPersistence{}).GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}
