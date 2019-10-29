package persistence
// repository という名前にしたいが domain 配下の repository とパッケージ名が被ってしまうため persistence で代替

import (
	"context"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/repository"
	"time"

	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
)

type bookPersistence struct {}

// NewBookPersistence : Book データに関する Persistence を生成
func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// GetAll : DB から Book データを全件取得（BookRepository インターフェースの GetAll() を実装したもの）
//  -> 本来は DB からデータを取得するが、簡略化のために省略（モックデータを返却）
func (bp bookPersistence) GetAll(context.Context) ([]*model.Book, error) {
	book1 := model.Book{}
	book1.Id = 1
	book1.Title = "Test1"
	book1.Author = "Tester1"
	book1.IssuedAt = time.Now()
	book1.CreatedAt = time.Now()
	book1.UpdatedAt = time.Now()

	book2 := model.Book{}
	book2.Id = 2
	book2.Title = "Test2"
	book2.Author = "Tester2"
	book2.IssuedAt = time.Now()
	book2.CreatedAt = time.Now()
	book2.UpdatedAt = time.Now()

	return []*model.Book{ &book1, &book2 }, nil
}
