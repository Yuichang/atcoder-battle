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

// jsonから取得するデータの構造体
type AtCHistory struct {
	IsRated           bool   `json:"IsRated"`
	Place             int    `json:"Place"`
	ContestScreenName string `json:"ContestScreenName"`
}

// HTMLに表示させるデータの構造体
type CompareResult struct {
	ContestShortName string
	WinUser          int
	User1Place       int
	User2Place       int
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

// 一旦二人のユーザーのデータを比較して返す
func CompareUsers(user1 string, user2 string) ([]CompareResult, error) {
	// 二人のユーザーデータを取得
	var user1History, user2History []AtCHistory
	var err error

	var compareResults []CompareResult

	user1History, err = GetUserData(user1)
	if err != nil {
		return nil, err
	}
	user2History, err = GetUserData(user2)
	if err != nil {
		return nil, err
	}

	// 共通のRatedコンテストを調べて、順位比較で勝ち負けを集計

	var user1Wins, user2Wins, draw, winUser int

	for i := 0; i < len(user1History); i++ {
		for j := 0; j < len(user2History); j++ {
			// 共通のRatedコンテストが見つかった
			if user1History[i].ContestScreenName == user2History[j].ContestScreenName && user1History[i].IsRated && user2History[j].IsRated {
				winUser = 0
				if user1History[i].Place < user2History[j].Place {
					user1Wins++
					winUser = 1
				} else if user1History[i].Place > user2History[j].Place {
					user2Wins++
					winUser = 2
				} else {
					draw++
					winUser = 0
				}

				// 初期の特殊なコンテスト名の時は別処理を後でする
				contestShortName := user1History[i].ContestScreenName[:6]
				compareResults = append(compareResults, CompareResult{
					ContestShortName: contestShortName,
					WinUser:          winUser,
					User1Place:       user1History[i].Place,
					User2Place:       user2History[j].Place,
				})
			}
		}
	}

	return compareResults, nil
}

// 取得した複数のユーザーのデータを処理
/*func CulcUserData([]AtCHistory... ) {
	// ここでユーザーデータを処理する
}*/
