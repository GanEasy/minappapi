package minappapi

import (
	"errors"
	"time"
)

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
	DB().Where(Subscribe{FansID: fans.ID, PostID: post.ID, Push: false}).Order("id desc").First(subscribe)

	// 7天前
	day7 := time.Now().AddDate(0, 0, -7)
	if subscribe.ID > 0 && day7.Before(subscribe.CreatedAt) { // 有订阅id且订阅时间在7天内
		return true
	}
	return false
}

func PostSubcribe(openID, formID, url string) {

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
