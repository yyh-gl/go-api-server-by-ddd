package usecase

import (
	"context"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/repository"
)

// BookUseCase : Book における UseCase のインターフェイス
type BookUseCase interface {
	GetAll(context.Context) ([]*model.Book, error)
}

type bookUseCase struct {
	bookPersistence repository.BookRepository
}

// NewBookUseCase : Book データに関する UseCase を生成
func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookPersistence: br,
	}
}

// GetAll : Book データを全件取得するためのユースケース
//  -> 本システムではあまりユースケース層の恩恵を受けれないが、もう少し大きなシステムになってくると、
//    「ドメインモデルの調節者」としての役割が見えてくる
func (bu bookUseCase) GetAll(ctx context.Context) (books []*model.Book, err error) {
	// Persistence（Repository）を呼出
	books, err = bu.bookPersistence.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
