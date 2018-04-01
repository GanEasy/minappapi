package main

import (
	"net/http"
	"strconv"
	"time"

	cpi "github.com/GanEasy/minappapi"
	"github.com/labstack/echo"
)

//CheckSubcribeUpdate  每天处理订阅更新
func CheckSubcribeUpdate() {
	ticker := time.NewTicker(time.Hour * 6)
	for _ = range ticker.C {
		go cpi.RunSubcribePostUpdateCheck()
	}
}
func main() {
	go CheckSubcribeUpdate()
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
	// 获取 url geturl
	e.GET("/geturl", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		ret := cpi.GetPostByID(int64(id))
		return c.JSON(http.StatusOK, ret)
	})
	// 获取 id getid
	e.GET("/getid", func(c echo.Context) error {
		url := c.QueryParam("url")
		id := cpi.GetIDByURL(url)
		return c.JSON(http.StatusOK, id)
	})
	// 检查订阅
	e.GET("/checksubscribe", func(c echo.Context) error {
		openID := c.QueryParam("openid")
		url := c.QueryParam("url")
		cs := cpi.CheckSubcribe(openID, url)
		type Ret struct {
			Status bool
		}
		return c.JSON(http.StatusOK, Ret{Status: cs})
	})
	// 取消订阅
	e.GET("/cancelsubscribe", func(c echo.Context) error {
		openID := c.QueryParam("openid")
		url := c.QueryParam("url")
		cs := cpi.CancelSubcribe(openID, url)
		type Ret struct {
			Status bool
		}
		return c.JSON(http.StatusOK, Ret{Status: cs})
	})
	// 反馈问题
	e.POST("/feedback", func(c echo.Context) error {
		// Feedback
		type Feedback struct {
			OpenID  string `json:"openid" form:"openid" query:"openid"`
			FormID  string `json:"formid" form:"formid" query:"formid"`
			Problem string `json:"problem" form:"problem" query:"problem"`
		}
		type Ret struct {
			Status bool
		}
		f := new(Feedback)
		if err := c.Bind(f); err != nil {
			return c.JSON(http.StatusOK, Ret{Status: false})
		}
		openID := f.OpenID
		// openID := c.QueryParam("openid")
		problem := f.Problem
		formID := f.FormID
		// openID := c.FormValue("openid")
		// // openID := c.QueryParam("openid")
		// answer := c.FormValue("answer")
		// formID := c.FormValue("formid")
		// answer := c.QueryParam("answer")
		cs := cpi.PostFeedback(openID, formID, problem)
		return c.JSON(http.StatusOK, Ret{Status: cs})
	})
	// 订阅
	e.GET("/subscribe", func(c echo.Context) error {
		openID := c.QueryParam("openid")
		url := c.QueryParam("url")
		formID := c.QueryParam("formid")
		cs := cpi.PostSubcribe(openID, formID, url)
		type Ret struct {
			Status bool
		}
		return c.JSON(http.StatusOK, Ret{Status: cs})
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
		ret, _ := cpi.GetBookMenu(urlStr)
		return c.JSON(http.StatusOK, ret)
	})
	// 记录分享 (我们现在通过分享次数进行排序)
	e.GET("/sharelog", func(c echo.Context) error {
		openID := c.QueryParam("openid")
		url := c.QueryParam("url")
		cs := cpi.ShareLog(openID, url)
		type Ret struct {
			Status bool
		}
		return c.JSON(http.StatusOK, Ret{Status: cs})

	})
	// 记录分享 (我们现在通过分享次数进行排序)
	e.GET("/sharerank", func(c echo.Context) error {
		posts := cpi.GetShareRank()
		return c.JSON(http.StatusOK, posts)

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

	// 获取小说目录正文
	e.GET("/getbookmenu", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		if urlStr == "" {
			return c.JSON(http.StatusOK, "0")
		}
		ret, _ := cpi.GetBookMenu(urlStr)
		return c.JSON(http.StatusOK, ret)
	})

	// 获取小说章节正文
	e.GET("/getbookcontent", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		if urlStr == "" {
			return c.JSON(http.StatusOK, "0")
		}
		ret, _ := cpi.GetBookContent(urlStr)
		return c.JSON(http.StatusOK, ret)
	})

	// 获取推荐码(图片资源)
	e.GET("/qrcode", func(c echo.Context) error {
		scene := c.QueryParam("scene")
		page := `pages/index/index`
		if scene == "" {
			return c.HTML(http.StatusOK, "")
		}
		fileName, err := cpi.GetwxCodeUnlimit(scene, page)
		if err == nil {
			http.ServeFile(c.Response().Writer, c.Request(), fileName)
		} else {
			http.ServeFile(c.Response().Writer, c.Request(), fileName)
		}
		var err2 error
		return err2
	})
	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8009"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
