package minappapi

import "testing"

func Test_CheckSubcribe(t *testing.T) {
	b2 := CheckSubcribe("oe9Hq0GwS4umXENTCn4lJgxaNVog", "http://book.zongheng.com/showchapter/523438.html")
	t.Fatal(b2)
}

func Test_PostSubcribe(t *testing.T) {
	b2 := PostSubcribe("oe9Hq0GwS4umXENTCn4lJgxaNVog", "xxss", "http://book.zongheng.com/showchapter/523438.html")
	t.Fatal(b2)
}
