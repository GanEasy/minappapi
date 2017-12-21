package main

import (
	"net/http"

	cpi "github.com/GanEasy/minappapi"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to reader minapp api!")
	})
	// 获取openid
	e.GET("/getopenid", func(c echo.Context) error {
		code := c.QueryParam("code")
		ret, _ := cpi.GetOpenID(code)
		return c.JSON(http.StatusOK, ret)
	})
	// 获取openid
	e.GET("/checksubscribe", func(c echo.Context) error {
		openID := c.QueryParam("open_id")
		url := c.QueryParam("url")
		cs := cpi.CheckSubcribe(openID, url)
		return c.JSON(http.StatusOK, cs)
	})
	// 获取openid
	e.GET("/subscribe", func(c echo.Context) error {
		openID := c.QueryParam("open_id")
		url := c.QueryParam("url")
		formID := c.QueryParam("form_id")
		cs := cpi.PostSubcribe(openID, formID, url)
		return c.JSON(http.StatusOK, cs)
	})
	// 解密数据内容
	e.GET("/crypt", func(c echo.Context) error {
		sessionKey := c.QueryParam("sk")
		encryptedData := c.QueryParam("ed")
		iv := c.QueryParam("iv")
		ret, _ := cpi.GetCryptData(sessionKey, encryptedData, iv)
		return c.JSON(http.StatusOK, ret)
	})
	// 获取列表
	e.GET("/getlist", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		if urlStr == "" {
			return c.JSON(http.StatusOK, "0")
		}
		ret, _ := cpi.GetList(urlStr)
		return c.JSON(http.StatusOK, ret)
	})
	// 获取正文
	e.GET("/getcontent", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		if urlStr == "" {
			return c.JSON(http.StatusOK, "0")
		}
		ret, _ := cpi.GetContent(urlStr)
		return c.JSON(http.StatusOK, ret)
	})
	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8009"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
