package minappapi

// GetPostByURL Post 如果没有的话进行初始化
func (post *Post) GetPostByURL(url string) {
	DB().Where(Post{URL: url}).FirstOrCreate(post)
}

// GetFansByOpenID Fans 如果没有的话进行初始化
func (fans *Fans) GetFansByOpenID(openID string) {
	DB().Where(Fans{OpenID: openID}).FirstOrCreate(fans)
}

// // GetFansByOpenID Fans 如果没有的话进行初始化
// func (post *Post) GetPostByURL(url string) {
// 	DB().Where(Post{URL: url}).FirstOrCreate(post)
// }

// Save Post
func (post *Post) Save() {
	DB().Save(&post)
}

// GetSubscribePost Post
func (post *Post) GetSubscribePost() []Post {
	var posts []Post
	// DB().Where(&Post{}).Find(&posts)
	DB().Where("fol_num > 0").Find(&posts)
	return posts
}

// Save Fans
func (fans *Fans) Save() {
	DB().Save(&fans)
}

// Save Subscribe
func (subscribe *Subscribe) Save() {
	DB().Save(&subscribe)
}

// GetSubscribeByPostID Subscribe
func (subscribe *Subscribe) GetSubscribeByPostID(postID uint) []Subscribe {
	var subscribes []Subscribe
	DB().Where(Subscribe{PostID: postID, Push: false}).Find(&subscribes)
	return subscribes
}

// GetSubscribeCheckList Subscribe
func (subscribe *Subscribe) GetSubscribeCheckList() []Subscribe {
	var subscribes []Subscribe
	DB().Find(&subscribes)
	return subscribes
}
