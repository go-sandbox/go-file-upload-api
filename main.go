package main

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

func main() {
	router := gin.Default()
	// ルートディレクトリは、assetsディレクトリ内のリソースをレスポンスとして返す
	router.Use(static.Serve("/", static.LocalFile("./assets", true)))

	// 画像アップロード
	router.POST("/upload", func(ctx *gin.Context) {
		// ファイル以外の場合エラーを返す
		img, err := imageupload.Process(ctx.Request, "file")
		if err != nil {
			panic(err)
		}

		// イメージを200x200に変換
		thumb, err := imageupload.ThumbnailPNG(img, 200, 200)
		if err != nil {
			panic(err)
		}
		// ハッシュを名前に指定して uploadディレクトリに保存
		h := sha1.Sum(thumb.Data)
		thumb.Save(fmt.Sprintf("upload/%s_%x.png",
			time.Now().Format("20060102150405"), h[:4]))
	})

	// ポート7777でサーバー起動
	router.Run(":7777")
}
