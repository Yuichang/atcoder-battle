package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*import(
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)*/

// 取得する構造体
type AtCHistory struct {
	IsRated           bool   `json:"IsRated"`
	Place             int    `json:"Place"`
	ContestScreenName string `json:"ContestScreenName"`
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

// 取得した複数のユーザーのデータを処理
/* func CulcUserData([]AtCHistory... ) {
	// ここでユーザーデータを処理する
}*/
