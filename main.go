package main

import (
	"fmt"
	"net/http"
)

// メイン処理
func main() {

	mux := http.NewServeMux()

	// ハンドラ
	mux.HandleFunc("/", index)
	mux.HandleFunc("/user", getUser)

	// サーバー設定
	server := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: mux,
	}
	server.ListenAndServe()
}

// ハンドラー
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
