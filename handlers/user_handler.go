package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ShowResult(c *gin.Context) {
	// 一旦ユーザーは二人だけにする（後日拡張)

	c.HTML(http.StatusOK, "result.html", gin.H{
		"user1": c.Query("user1"),
		"user2": c.Query("user2"),
	})
}

// サンプルページを表示するだけ
func ShowSample(c *gin.Context) {
	c.HTML(http.StatusOK, "sample.html", nil)
}
