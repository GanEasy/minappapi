package minappapi

import "testing"

func Test_CheckSubcribe(t *testing.T) {
	b2 := CheckSubcribe("oe9Hq0GwS4umXENTCn4lJgxaNVog", "http://book.zongheng.com/showchapter/523438.html")
	t.Fatal(b2)
}
