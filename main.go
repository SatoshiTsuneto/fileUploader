package main

import (
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

// 取得したファイルを保存する関数
func fileSave(c echo.Context) error {
	// フォームから送られてきたファイルを取得
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// 開く
	data, err := file.Open()
	if err != nil {
		return err
	}
	defer data.Close()

	// ファイルの作成
	dst, err := os.Create("./file/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 取得したファイルの保存
	_, err = io.Copy(dst, data)
	if err != nil {
		return err
	}

	// リザルトの送信
	return c.HTML(http.StatusOK, "<b>保存完了！</b>")
}


func main() {
	e := echo.New()
	e.POST("/file", fileSave)
	e.Logger.Fatal(e.Start(":9999"))
}
