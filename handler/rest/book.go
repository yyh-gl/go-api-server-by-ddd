package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
	"github.com/yyh-gl/go-api-server-by-ddd/usecase"
)

// BookHandler : Book における Handler のインターフェイス
type BookHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type bookHandler struct {
	bookUseCase usecase.BookUseCase
}

// NewBookUseCase : Book データに関する Handler を生成
func NewBookHandler(bu usecase.BookUseCase) BookHandler {
	return &bookHandler{
		bookUseCase: bu,
	}
}

// BookIndex : GET /books -> Book データ一覧を返す
func (bh bookHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	var books []*model.Book
	var err error

	ctx := r.Context()

	// ユースケースの呼出
	books, err = bh.bookUseCase.GetAll(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(books); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
