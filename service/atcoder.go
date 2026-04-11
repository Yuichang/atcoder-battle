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
// 詳細情報表示するときだけ使う形式に変えるかも。
type CompareDetail struct {
	ContestShortName string
	MatchWinner      int
	User1Place       int
	User2Place       int
}

// バトルの結果を格納する構造体
type BattleResult struct {
	User1Name string
	User2Name string
	User1Wins int
	User2Wins int
	Winner    string
	Draw      int
	Total     int
	Results   []CompareDetail
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
func CompareUsers(user1 string, user2 string) (BattleResult, error) {
	// 二人のユーザーデータを取得
	var user1History, user2History []AtCHistory
	var err error

	// バトルの詳細結果を格納する
	var compareResults []CompareDetail

	user1History, err = GetUserData(user1)
	if err != nil {
		return BattleResult{}, err
	}
	user2History, err = GetUserData(user2)
	if err != nil {
		return BattleResult{}, err
	}

	// 共通のRatedコンテストを調べて、順位比較で勝ち負けを集計

	var user1Wins, user2Wins, draw, matchWinner int
	var winner string

	// map使って高速化させる（後でやる）
	for i := 0; i < len(user1History); i++ {
		for j := 0; j < len(user2History); j++ {
			// 共通のRatedコンテストが見つかった
			if user1History[i].ContestScreenName == user2History[j].ContestScreenName && user1History[i].IsRated && user2History[j].IsRated {
				matchWinner = 0
				if user1History[i].Place < user2History[j].Place {
					user1Wins++
					matchWinner = 1
				} else if user1History[i].Place > user2History[j].Place {
					user2Wins++
					matchWinner = 2
				} else {
					draw++
					matchWinner = 0
				}

				// 初期の特殊なコンテスト名の時は別処理を後でする
				contestShortName := user1History[i].ContestScreenName[:6]

				compareResults = append(compareResults, CompareDetail{
					ContestShortName: contestShortName,
					MatchWinner:      matchWinner,
					User1Place:       user1History[i].Place,
					User2Place:       user2History[j].Place,
				})
			}
		}
	}

	// 勝者を決定させる
	if user1Wins > user2Wins {
		winner = user1
	} else if user1Wins < user2Wins {
		winner = user2
	} else {
		winner = "Draw"
	}

	// バトルの結果をまとめる
	battleResult := BattleResult{
		User1Name: user1,
		User2Name: user2,
		User1Wins: user1Wins,
		User2Wins: user2Wins,
		Winner:    winner,
		Draw:      draw,
		Total:     user1Wins + user2Wins + draw,
		Results:   compareResults,
	}

	return battleResult, nil
}

// 取得した複数のユーザーのデータを処理
/*func CulcUserData([]AtCHistory... ) {
	// ここでユーザーデータを処理する
}*/
