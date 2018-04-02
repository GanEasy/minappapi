package minappapi

import "testing"

func Test_GetContent(t *testing.T) {
	// a2, _ := GetContent("http://book.zongheng.com/showchapter/730066.html")
	a2, _ := GetContent("http://www.biquge5.com/0_9/5018.html")
	t.Fatal(a2)
}
