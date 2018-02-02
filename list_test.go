package minappapi

import "testing"

func Test_GetList(t *testing.T) {
	// l2, _ := GetList("http://book.zongheng.com/showchapter/523438.html")
	// l2, _ := GetList("http://www.3dllc.cc/html/86/86244/")
	l2, _ := GetList("http://longfu8.com/")
	t.Fatal(l2)
}
