package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// スクレイピングする構造体
type AtCHistory struct {
	IsRated           bool   `json:"IsRated"`
	Place             int    `json:"Place"`
	ContestScreenName string `json:"ContestScreenName"` //(abc242的な)
}

// ユーザー名を取得して返す
func GetUserData(username string) ([]AtCHistory, error) {
	url := "https://atcoder.jp/users/" + username + "/history/json"

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var history []AtCHistory

	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		return nil, err
	}

	fmt.Println(history)

	return history, nil
}

/*func main() {
	// 実行テスト
	username := "tourist"
	_, err := GetUserData(username)
	if err != nil {
		fmt.Println("err")
	}
}*/
