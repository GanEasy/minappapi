package minappapi

import (
	"errors"
	"fmt"

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
