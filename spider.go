package minappapi

import (
	"encoding/json"

	"github.com/yizenghui/reader"
)

// RunSubcribePostUpdateCheck 检查订阅状况并向订阅者推送更新
func RunSubcribePostUpdateCheck() {
	var post Post
	posts := post.GetSubscribePost()
	if len(posts) > 0 {
		for _, post := range posts {
			CheckPostChapterUpdateAndPushSubscribe(&post)
		}
	}
}

// CheckPostChapterUpdateAndPushSubscribe 获取文章更新并推送通知
func CheckPostChapterUpdateAndPushSubscribe(post *Post) bool {

	b, err := GetPostChapterByte(post.URL)
	if err != nil {
		return false
	}

	if post.ChapterFragments != string(b) {
		post.ChapterFragments = string(b)
		// todo 发通知咯 有变化了
		NoticeSubscribePostUpdate(post)

		return true
	}

	return false
}

// CheckPostChapterUpdateAndPushSubscribeByURL 获取文章更新并推送通知
func CheckPostChapterUpdateAndPushSubscribeByURL(url string) bool {
	b, err := GetPostChapterByte(url)
	if err != nil {
		return false
	}

	post, err := GetPostByURL(url)
	if err != nil {
		return false
	}

	if post.ChapterFragments != string(b) {
		post.ChapterFragments = string(b)
		// todo 发通知咯 有变化了
		NoticeSubscribePostUpdate(post)

		return true
	}

	return false
}

// GetPostChapterByte 获取链接内容
func GetPostChapterByte(url string) (b []byte, err error) {
	data, err := GetList(url)
	if err != nil {
		return
	}
	list := GetPostChapter(data.Links)
	b, err = json.Marshal(list)
	return
}

// GetPostChapter 获取章节片段
func GetPostChapter(links []reader.Link) []reader.Link {

	var list []reader.Link
	if len(links) > 10 {
		for _, v := range links[0:5] {
			list = append(list, v)
		}

		for _, v := range links[len(links)-5:] {
			list = append(list, v)
		}
	}
	return list
}
