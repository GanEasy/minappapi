package minappapi

import (
	"errors"
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
	DB().Where(&Subscribe{FansID: fans.ID, PostID: post.ID, Push: false}).Order("id desc").First(&subscribe)

	// 7天前
	day7 := time.Now().AddDate(0, 0, -7)
	if subscribe.ID > 0 && day7.Before(subscribe.CreatedAt) { // 有订阅id且订阅时间在7天内
		return true
	}
	return false
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
	subscribe := Subscribe{FansID: fans.ID, PostID: post.ID, FormID: formID, Push: false}

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
func NoticeSubscribePostUpdate(url string) bool {
	post, err := GetPostByURL(url)
	if err != nil {
		return false
	}
	var subscribe Subscribe
	subscribes := subscribe.GetSubscribeByPostID(post.ID)
	if len(subscribes) > 0 {
		// for _, sub := range subscribes {

		// }
	}
	return true
}
