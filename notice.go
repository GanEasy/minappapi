package minappapi

// NoticeFansPostUpdate 通知关注更新书籍已经更新
func NoticeFansPostUpdate(url string) bool {
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
