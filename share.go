package minappapi

import "time"

//ShareLog 检查订阅状况
func ShareLog(openID, url string) bool {
	fans, err := GetFansByOpenID(openID)
	if err != nil {
		return false
	}
	post, err := GetPostByURL(url)
	if err != nil {
		return false
	}

	var share Share
	// 这里面有bug无法在结构体中查询false
	DB().Where(&Share{FansID: fans.ID, PostID: post.ID}).First(&share)
	// 已经分享过了
	if share.ID > 0 {
		share.SubNum++
		// 限制1天内没分享过 不允许分享刷榜
		day1 := time.Now().AddDate(0, 0, -1)
		if share.UpdatedAt.Before(day1) {
			post.ShareNum++
			post.Save()
		}
	} else {
		// 首次分享
		share.SubNum = 1
		post.ShareNum++
		post.Save()
	}
	share.Save()

	return false
}

//GetShareRank 获取分享排行榜
func GetShareRank() []Post {
	var post Post
	posts := post.GetShareRankPost()
	return posts
}
