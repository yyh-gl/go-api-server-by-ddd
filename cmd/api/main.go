package main

import (
	"fmt"
	"github.com/yyh-gl/go-api-server-by-ddd/infra/persistence"
	"github.com/yyh-gl/go-api-server-by-ddd/usecase"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yyh-gl/go-api-server-by-ddd/handler"
)

func main() {
	// 依存関係を注入
	// DI ライブラリを使えば、もっとスマートになるはず
	bookPersistence := persistence.NewBookPersistence()
	bookUseCase := usecase.NewBookUseCase(bookPersistence)
	bookHandler := handler.NewBookHandler(bookUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/books", bookHandler.Index)

	// サーバ起動
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:3000")
	fmt.Println("========================")
	log.Fatal(http.ListenAndServe(":3000", router))
}
