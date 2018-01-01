package minappapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

//TokenServe token 服务器
var TokenServe *DefaultAccessTokenServer

func init() {

	TokenServe = NewDefaultAccessTokenServer(config.ReaderMinApp.AppID, config.ReaderMinApp.AppSecret)

}

//GetSubcribePost 检查订阅状况
// func GetSubcribePost() {
// 	var post Post
// 	posts := post.GetSubscribePost()
// }

//CheckSubcribe 检查订阅状况
func CheckSubcribe(openID, url string) bool {
	fans, err := GetFansByOpenID(openID)
	if err != nil {
		return false
	}
	post, err := GetPostByURL(url)
	if err != nil {
		return false
	}

	var subscribe Subscribe
	// 这里面有bug无法在结构体中查询false
	DB().Where(&Subscribe{FansID: fans.ID, PostID: post.ID}).Where("push = ?", false).Order("id desc").First(&subscribe)

	// 7天前
	day7 := time.Now().AddDate(0, 0, -7)
	if subscribe.ID > 0 && day7.Before(subscribe.CreatedAt) { // 有订阅id且订阅时间在7天内
		// log.Print(subscribe.ID)
		return true
	}
	return false
}

//CancelSubcribe 取消订阅
func CancelSubcribe(openID, url string) bool {
	fans, err := GetFansByOpenID(openID)
	if err != nil {
		return false
	}
	post, err := GetPostByURL(url)
	if err != nil {
		return false
	}
	// 订阅软删除
	DB().Where(&Subscribe{FansID: fans.ID, PostID: post.ID}).Where("push = ?", false).Delete(Subscribe{})

	return true
}

// PostSubcribe 提交关注请求
func PostSubcribe(openID, formID, url string) bool {
	if formID == "" {
		return false
	}
	fans, err := GetFansByOpenID(openID)
	if err != nil {
		return false
	}
	post, err := GetPostByURL(url)
	if err != nil {
		return false
	}
	// 如果没有更新过章节内容，关注时先初始化
	if post.ChapterFragments == "" {
		b, err := json.Marshal(GetPostChapter(url))
		if err != nil {
			return false
		}
		post.ChapterFragments = string(b)
	}
	// 增加关注..
	post.FolNum++
	post.SubNum++
	post.Save()

	subscribe := Subscribe{FansID: fans.ID, PostID: post.ID, FormID: formID, Push: false, OpenID: openID}

	DB().Create(&subscribe)
	if subscribe.ID > 0 {
		return true
	}
	return false
}

// GetPostByURL 获取post
func GetPostByURL(url string) (*Post, error) {
	var err error
	var post Post
	if url != "" {
		post.GetPostByURL(url)
	} else {
		err = errors.New(string(`url is empty!!!`))
	}
	return &post, err
}

// NoticeSubscribePostUpdate 通知关注更新书籍已经更新
func NoticeSubscribePostUpdate(post *Post) bool {
	var subscribe Subscribe
	subscribes := subscribe.GetSubscribeByPostID(post.ID)
	if len(subscribes) > 0 {
		for _, sub := range subscribes {
			//
			link := fmt.Sprintf("pages/list/index?url=%v", url.QueryEscape(post.URL))
			SendPostUpdateMSG(sub.OpenID, sub.FormID, post.Title, "您所关注的目录已更新！", link)
			sub.Push = true
			sub.Save()
		}
	}
	post.FolNum = 0
	post.Save()
	return true
}
