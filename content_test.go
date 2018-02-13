package minappapi

import "testing"

func Test_GetContent(t *testing.T) {
	a2, _ := GetContent("http://book.zongheng.com/showchapter/730066.html")
	t.Fatal(a2)
}
