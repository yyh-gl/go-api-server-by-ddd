package rest

// Handler 層を変えるだけで、例えば CLI にも簡単に対応可能

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
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
	// request : 本 API のリクエストパラメータ
	//  -> こんな感じでリクエストも受け取れますが、今回は使いません
	type request struct {
		Begin uint `query:"begin"`
		Limit uint `query:"limit"`
	}

	// bookField : response 内で使用する Book を表す構造体
	//  -> ドメインモデルの Book に HTTP の関心事である JSON タグを付与したくないために Handler 層で用意
	//     簡略化のために JSON タグを付与したドメインモデルを流用するプロジェクトもしばしば見かける
	type bookField struct {
		Id       int64     `json:"id"`
		Title    string    `json:"title"`
		Author   string    `json:"author"`
		IssuedAt time.Time `json:"issued_at"`
	}

	// response : 本 API のレスポンス
	type response struct {
		Books []bookField `json:"books"`
	}

	ctx := r.Context()

	// ユースケースの呼出
	books, err := bh.bookUseCase.GetAll(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	res := new(response)
	for _, book := range books {
		var bf bookField
		bf = bookField(*book)
		res.Books = append(res.Books, bf)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
