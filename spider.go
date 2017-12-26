package minappapi

import "github.com/yizenghui/reader"

//GetSubcribePost 检查订阅状况
// func GetSubcribePost() {
// 	var post Post
// 	posts := post.GetSubscribePost()
// }

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
