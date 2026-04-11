package handlers

import (
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
			"message": "ユーザーデータの取得に失敗しました",
		})
		return
	}

	c.HTML(http.StatusOK, "result.html", gin.H{
		/*"contestShortName": result.ContestShortName,
		"winUser":          result.WinUser,
		"user1Place":       result.User1Place,
		"user2Place":       result.User2Place,*/
		"results": result,
	})
}

/*func ShowResult(c *gin.Context) {
	// 一旦ユーザーは二人だけにする

}*/

// サンプルページを表示するだけ
func ShowSample(c *gin.Context) {
	c.HTML(http.StatusOK, "sample.html", nil)
}
