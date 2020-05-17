package main

import (
	"encoding/json"
	"net/http"
)

// UserResponse レスポンスデータ
type UserResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Introduce    string `json:"introduce"`
	ErrorMessage string `json:"errorMessage"`
}

// ユーザー取得
func getUser(w http.ResponseWriter, r *http.Request) {

	// ステータスを取得
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// レスポンスデータ生成
	userInfo := UserResponse{1, "1", "aaa@aaa", "自己紹介", ""}

	js, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)

}
