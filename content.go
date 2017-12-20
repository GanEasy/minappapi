package minappapi

import (
	"fmt"
	"html/template"

	"github.com/lunny/html2md"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/yizenghui/reader"
)

//ContentInfo 数据包
type ContentInfo struct {
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
	PubAt   string        `json:"pub_at"`
}

//GetContent 获取正文
func GetContent(urlStr string) (ci ContentInfo, err error) {

	info, err := reader.GetContent(urlStr)
	if err != nil {
		return ci, err
	}

	md := html2md.Convert(info.Content)
	input := []byte(md)
	unsafe := blackfriday.MarkdownCommon(input)
	content := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	info.Content = fmt.Sprintf(`%v`, string(content[:]))

	// // 给图片加上 最大宽度
	// info.Content = strings.Replace(info.Content, `<img src=`, `<img style="max-width:100%" src=`, -1)
	// info.Content = strings.Replace(info.Content, `<section>`, `<div>`, -1)
	// info.Content = strings.Replace(info.Content, `</section>`, `</div>`, -1)

	return ContentInfo{
		info.Title,
		template.HTML(info.Content),
		info.PubAt,
	}, nil
}
