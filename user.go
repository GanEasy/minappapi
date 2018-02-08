package minappapi

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/chanxuehong/wechat.v2/mp/message/template"
	wxbizdatacrypt "github.com/yilee/wx-biz-data-crypt"
)

// OpenIDData 开放数据 openID
type OpenIDData struct {
	ErrCode int64  `json:"errcode"`
	OpenID  string `json:"openid"`
}

//GetOpenID 获取微信小程序上报的openid 此ID暂不加密处理
func GetOpenID(code string) (OpenIDData, error) {
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

	HTTPGetJSON(url, &ret)
	var err error

	if ret.ErrCode != 0 {
		err = errors.New(string(ret.ErrCode))
	}

	return OpenIDData{ret.ErrCode, ret.OpenID}, err
}

//SendPostUpdateMSG 发送更新通知
func SendPostUpdateMSG(openID, formID, title, page string) error {
	//
	type Ret struct {
		ErrCode int64  `json:"errcode"`
		ErrMSG  string `json:"errmsg"`
	}
	var ret Ret

	type DataItem struct {
		Value string `json:"value"`
		Color string `json:"color"`
	}

	type Template struct {
		Touser          string      `json:"touser"`
		TemplateID      string      `json:"template_id"`
		Page            string      `json:"page"`
		FormID          string      `json:"form_id"`
		Data            interface{} `json:"data"`
		EmphasisKeyword string      `json:"emphasis_keyword"`
	}

	//MSG 关注通知消息结构
	type MSG struct {
		Title    template.DataItem `json:"keyword1"`
		CATEGORY template.DataItem `json:"keyword2"`
	}

	data := Template{
		Touser:     openID,
		TemplateID: "QEhBZIivAI5x0hbWEp4IqMKAb3RhLXCl3eBr1GC_7FE",
		Page:       page,
		FormID:     formID,
		Data: MSG{
			Title:    template.DataItem{Value: title, Color: ""},
			CATEGORY: template.DataItem{Value: "文章", Color: ""},
		},
		EmphasisKeyword: "",
	}

	token, err2 := TokenServe.Token()
	if err2 != nil {

		return err2
	}
	url := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%v`, token)

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	HTTPPostJSON(url, b, &ret)

	if ret.ErrCode != 0 {
		err = errors.New(string(ret.ErrCode))
	}

	return err
}

//GetwxCodeUnlimit 发送更新通知
func GetwxCodeUnlimit(scene, page string) (file string, err error) {

	name := GetMd5String(fmt.Sprintf(`%v%v`, scene, page))
	file = fmt.Sprintf(`file/%v.jpg`, name)

	_, err2 := os.Stat(file)
	if os.IsNotExist(err2) {

		type Template struct {
			Scene     string      `json:"scene"`
			Page      string      `json:"page"`
			Width     int         `json:"width"`
			AutoColor bool        `json:"auto_color"`
			LineColor interface{} `json:"line_color"`
		}

		data := Template{
			Scene: scene,
			Page:  page,
		}

		token, err2 := TokenServe.Token()
		if err2 != nil {
			return "", err2
		}
		url := fmt.Sprintf(`https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%v`, token)

		b, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		_, err = SaveQrcodeImg(url, file, b)
	}
	return file, err
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SaveQrcodeImg 保存图片到本地
func SaveQrcodeImg(imageURL, saveName string, body []byte) (n int64, err error) {
	out, err := os.Create(saveName)
	defer out.Close()
	if err != nil {
		return
	}
	// text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
	// application/json; charset=utf-8
	resp, err := httpClient.Post(imageURL, "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8", bytes.NewReader(body))

	if err != nil {
		return
	}
	pix, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err != nil {
		return
	}
	n, err = io.Copy(out, bytes.NewReader(pix))

	if err != nil {
		return
	}
	// todo 获取图片类型
	// fmt.Println(resp.Header.Get("Content-Type"))
	return
}

// GetCryptData 解密数据
func GetCryptData(sessionKey, encryptedData, iv string) (*wxbizdatacrypt.UserInfo, error) {
	pc := wxbizdatacrypt.NewWXBizDataCrypt(config.ReaderMinApp.AppID, sessionKey)
	return pc.Decrypt(encryptedData, iv)
}

// GetFansByOpenID 解密数据
func GetFansByOpenID(openID string) (*Fans, error) {
	var err error
	var fans Fans
	if openID != "" {
		fans.GetFansByOpenID(openID)
	} else {
		err = errors.New(string(`openID is empty!!!`))
	}
	return &fans, err
}
