package handlers

import (
	"encoding/json"
	"net/http"

	//"github.com/Yuichang/atcoder-battle/models"
	"github.com/Yuichang/atcoder-battle/service"
	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ShowResult(c *gin.Context) {
	// 一旦ユーザーは二人だけにする（後日拡張)

	user1 := c.Query("user1")
	user2 := c.Query("user2")

	// ユーザーデータを取得後、比較してデータを返す
	result, err := service.CompareUsers(user1, user2)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "result.html", result)
}

// サンプルページを表示するだけ
func ShowSample(c *gin.Context) {
	c.HTML(http.StatusOK, "sample.html", nil)
}

// ユーザーが存在するか否かを確認
func CheckUser(c *gin.Context) {
	username := c.Query("username")

	url := "https://atcoder.jp/users/" + username + "/history/json"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(400, gin.H{"ok": false})
		return
	}
	defer resp.Body.Close()

	var history []map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		c.JSON(400, gin.H{"ok": false})
		return
	}

	if len(history) == 0 {
		c.JSON(400, gin.H{"ok": false})
		return
	}

	c.JSON(200, gin.H{"ok": true})
}
