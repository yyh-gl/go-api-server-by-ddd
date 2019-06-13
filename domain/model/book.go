package model

import "time"

type Book struct {
	Id       int64      `json:"id"`
	Title    string     `json:"title"`
	Author   string     `json:"author"`
	IssuedAt  time.Time `json:"issued_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
