package minappapi

// GetPostByURL Post 如果没有的话进行初始化
func (post *Post) GetPostByURL(url string) {
	DB().Where(Post{URL: url}).FirstOrCreate(&post)
}

// GetPostByID Post 如果没有的话进行初始化
func (post *Post) GetPostByID(id int64) {
	DB().First(&post, id)
}

// GetFansByOpenID Fans 如果没有的话进行初始化
func (fans *Fans) GetFansByOpenID(openID string) {
	DB().Where(Fans{OpenID: openID}).FirstOrCreate(&fans)
}

// // GetFansByOpenID Fans 如果没有的话进行初始化
// func (post *Post) GetPostByURL(url string) {
// 	DB().Where(Post{URL: url}).FirstOrCreate(post)
// }

// Save Post
func (post *Post) Save() {
	DB().Save(&post)
}

// Save Share
func (share *Share) Save() {
	DB().Save(&share)
}

// Save Feedback
func (feedback *Feedback) Save() {
	DB().Save(&feedback)
}

// GetSubscribePost Post
func (post *Post) GetSubscribePost() []Post {
	var posts []Post
	// DB().Where(&Post{}).Find(&posts)
	DB().Where("sub_num > 0").Find(&posts)
	// DB().Where("fol_num > 0").Find(&posts)
	return posts
}

// GetShareRankPost Post
func (post *Post) GetShareRankPost() []Post {
	var posts []Post
	// DB().Where(&Post{}).Find(&posts)
	// DB().Where("sub_num > 0").Where("share > 0").Order("share desc,id desc").Limit(100).Find(&posts)
	DB().Where("fol_num > 10").Where("share_num > 10").Order("share_num desc, id desc").Limit(100).Find(&posts)
	// DB().Where("fol_num > 0").Find(&posts)
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
