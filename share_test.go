package minappapi

import (
	"testing"
)

// func Test_MapMake(t *testing.T) {
// 	list, _ := GetList(`http://www.loxow.com/97371.shtml`)
// 	t.Fatal(list.Links[0:5])
// }

func Test_ShareLog(t *testing.T) {
	b := ShareLog(`oe9Hq0GwS4umXENTCn4lJgxaNVog`, `http://www.88dushu.com/xiaoshuo/0/511/`)
	t.Fatal(b)
}

func Test_GetShareRank(t *testing.T) {
	b := GetShareRank()
	t.Fatal(b)
}
