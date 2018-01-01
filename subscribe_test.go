package minappapi

import (
	"net/url"
	"testing"
)

func Test_CheckSubcribe(t *testing.T) {
	b2 := CheckSubcribe("oe9Hq0GwS4umXENTCn4lJgxaNVog", "http://longfu8.com/")
	t.Fatal(b2)
}

func Test_CancelSubcribe(t *testing.T) {
	b2 := CancelSubcribe("oe9Hq0GwS4umXENTCn4lJgxaNVog", "http://longfu8.com/")
	t.Fatal(b2)
}
func Test_PostSubcribe(t *testing.T) {
	b2 := PostSubcribe("oe9Hq0GwS4umXENTCn4lJgxaNVog", "xxss", "http://book.zongheng.com/showchapter/523438.html")
	t.Fatal(b2)
}

func Test_CheckPostChapterUpdateAndPushSubscribe(t *testing.T) {
	b2 := CheckPostChapterUpdateAndPushSubscribe("http://www.51shucheng.net/kehuan/santi/")
	t.Fatal(b2)
}

func Test_Token(t *testing.T) {

	t.Fatal(TokenServe.Token())
}

func Test_EnCodeURL(t *testing.T) {

	var urlStr string = "http://baidu.com/index.php/?abc=1_羽毛"

	l3, _ := url.Parse(urlStr)
	// t.Fatal(l3.Encode)
	t.Fatal(l3.Query().Encode())
}
