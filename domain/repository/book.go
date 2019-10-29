package repository

import "github.com/yyh-gl/go-api-server-by-ddd/domain/model"

// BookRepository : Book における Repository のインターフェイス
//  -> 依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
type BookRepository interface {
	GetAll() ([]*model.Book, error)
}
