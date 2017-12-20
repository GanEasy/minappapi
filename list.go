package minappapi

import "github.com/yizenghui/reader"

//GetList 获取链接列表
func GetList(urlStr string) (reader.Data, error) {
	return reader.GetList(urlStr)
}
