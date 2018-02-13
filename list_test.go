package minappapi

import "testing"

func Test_GetList(t *testing.T) {
	// l2, _ := GetList("http://book.zongheng.com/showchapter/523438.html")
	// l2, _ := GetList("http://www.3dllc.cc/html/86/86244/")
	// l2, _ := GetList("http://longfu8.com/")
	l2, _ := GetList("http://book.zongheng.com/showchapter/730066.html")
	t.Fatal(l2)
}

func Test_GetBookMenu(t *testing.T) {
	// l2, _ := GetList("http://book.zongheng.com/showchapter/523438.html")
	// l2, _ := GetList("http://www.3dllc.cc/html/86/86244/")
	// l2, _ := GetList("http://longfu8.com/")
	l2, _ := GetBookMenu("http://book.zongheng.com/showchapter/730066.html")
	t.Fatal(l2)
}
