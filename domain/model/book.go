package model

import "time"

// Book : Book を表すドメインモデル
// !! 重度のドメイン貧血症です !!
type Book struct {
	Id       int64
	Title    string
	Author   string
	IssuedAt time.Time
}
