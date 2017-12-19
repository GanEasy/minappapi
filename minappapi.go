package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	wxbizdatacrypt "github.com/yilee/wx-biz-data-crypt"
	"github.com/yizenghui/reader"
)

type tomlConfig struct {
	ReaderMinApp ReaderMinApp
}

//ReaderMinApp 配置
type ReaderMinApp struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

var config tomlConfig

//DefaultHttpClient 默认http.Client
var DefaultHttpClient *http.Client

func init() {
	client := *http.DefaultClient
	client.Timeout = time.Second * 5
	DefaultHttpClient = &client
	if _, err := toml.DecodeFile("conf.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}

//GetOpenID 获取微信小程序上报的openid 此ID暂不加密处理
func GetOpenID(c echo.Context) error {
	code := c.QueryParam("code")
	//
	type Ret struct {
		ErrCode    int64  `json:"errcode"`
		ErrMSG     string `json:"errmsg"`
		SessionKey string `json:"session_key"`
		ExpiresIn  int64  `json:"expires_in"`
		OpenID     string `json:"openid"`
	}
	var ret Ret

	url := fmt.Sprintf(`https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code`,
		config.ReaderMinApp.AppID,
		config.ReaderMinApp.AppSecret,
		code,
	)
	// fmt.Println(url)
	httpGetJSON(url, &ret)

	type RetData struct {
		ErrCode int64  `json:"errcode"`
		OpenID  string `json:"openid"`
	}
	return c.JSON(http.StatusOK, RetData{ret.ErrCode, ret.OpenID})
}

//GetContent 获取正文
func GetContent(c echo.Context) error {

	urlStr := c.QueryParam("url")

	info, err := reader.GetContent(urlStr)
	if err != nil {
		return c.String(http.StatusOK, "0")
	}

	input := []byte(info.Content)
	unsafe := blackfriday.MarkdownCommon(input)
	content := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// html := fmt.Sprintf(`<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	// 						<link rel="preload" href="https://yize.gitlab.io/css/main.css" as="style" />
	// 						%v`, string(content[:]))
	// return c.HTML(http.StatusOK, html)
	info.Content = fmt.Sprintf(`%v`, string(content[:]))

	// 给图片加上 最大宽度
	info.Content = strings.Replace(info.Content, `<img src=`, `<img style="max-width:100%" src=`, -1)
	info.Content = strings.Replace(info.Content, `<section>`, `<div>`, -1)
	info.Content = strings.Replace(info.Content, `</section>`, `</div>`, -1)

	type Info struct {
		Title   string        `json:"title"`
		Content template.HTML `json:"content"`
		PubAt   string        `json:"pub_at"`
	}

	return c.JSON(http.StatusOK, Info{
		info.Title,
		template.HTML(info.Content),
		info.PubAt,
	})
}

//GetList 获取列表 临时放在这里面，做小程序测试api
func GetList(c echo.Context) error {
	urlStr := c.QueryParam("url")
	if urlStr == "" {
		return c.JSON(http.StatusOK, "0")
	}
	links, _ := reader.GetList(urlStr)
	return c.JSON(http.StatusOK, links)
}

// Crypt 解密数据
func Crypt(c echo.Context) error {
	appID := config.ReaderMinApp.AppID
	sessionKey := c.QueryParam("sk")
	encryptedData := c.QueryParam("ed")
	iv := c.QueryParam("iv")
	pc := wxbizdatacrypt.NewWXBizDataCrypt(appID, sessionKey)
	userInfo, _ := pc.Decrypt(encryptedData, iv)
	return c.JSON(http.StatusOK, userInfo)
}

func httpGetJSON(url string, response interface{}) error {
	httpResp, err := DefaultHttpClient.Get(url)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return DecodeJSONHttpResponse(httpResp.Body, response)
}

//DecodeJSONHttpResponse 解决json
func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	// body2 := body
	// fmt.Println(body)
	// buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	// if err := json.Indent(buf, body2, "", "    "); err == nil {
	// 	body2 = buf.Bytes()
	// }
	// log.Printf("[WECHAT_DEBUG] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to reader minapp api!")
	})
	e.GET("/getopenid", GetOpenID)
	e.GET("/crypt", Crypt)
	e.GET("/getlist", GetList)
	e.GET("/getcontent", GetContent)
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8009"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
