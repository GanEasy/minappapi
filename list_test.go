package minappapi

import "testing"

func Test_GetList(t *testing.T) {
	l2, _ := GetList("http://book.zongheng.com/showchapter/523438.html")
	t.Fatal(l2)
}
