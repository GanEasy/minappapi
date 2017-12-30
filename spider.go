package minappapi

import (
	"encoding/json"

	"github.com/yizenghui/reader"
)

//GetSubcribePost 检查订阅状况
// func GetSubcribePost() {
// 	var post Post
// 	posts := post.GetSubscribePost()
// }

// CheckPostChapterUpdateAndPushSubscribe 获取文章更新并推送通知
func CheckPostChapterUpdateAndPushSubscribe(url string) bool {
	list := GetPostChapter(url)
	b, err := json.Marshal(list)
	if err != nil {
		return false
	}

	post, err := GetPostByURL(url)
	if err != nil {
		return false
	}

	if post.ChapterFragments != string(b) {
		// todo 发通知咯 有变化了
		
		return true
	}

	return false
}

// GetPostChapterByte 获取链接内容
func GetPostChapterByte(url string) (b []byte, err error) {
	list := GetPostChapter(url)
	b, err = json.Marshal(list)
	return
}

// GetPostChapter 获取章节片段
func GetPostChapter(url string) []reader.Link {
	list, err := GetList(url)
	if err != nil {

	}
	var links []reader.Link
	if len(list.Links) > 10 {
		for _, v := range list.Links[0:5] {
			links = append(links, v)
		}

		for _, v := range list.Links[len(list.Links)-5:] {
			links = append(links, v)
		}
	}
	return links
}
