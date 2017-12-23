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

// Save Fans
func (fans *Fans) Save() {
	DB().Save(&fans)
}

// Save Subscribe
func (subscribe *Subscribe) Save() {
	DB().Save(&subscribe)
}