package main

import (
	"github.com/Yuichang/atcoder-battle/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静的ファイル
	r.Static("/static", "./static")

	// テンプレート読み込み
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.ShowIndex)

	r.GET("/result", handlers.ShowResult)

	r.GET("/sample", handlers.ShowSample)

	r.GET("/check_user", handlers.CheckUser)

	// サーバ起動(http://localhost:8080)
	r.Run()
}
