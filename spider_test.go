package minappapi

import "testing"

// func Test_MapMake(t *testing.T) {
// 	list, _ := GetList(`http://www.loxow.com/97371.shtml`)
// 	t.Fatal(list.Links[0:5])
// }

func Test_GetPostChapter(t *testing.T) {
	list := GetPostChapter(`http://book.zongheng.com/showchapter/523438.html`)
	t.Fatal(list)
}
