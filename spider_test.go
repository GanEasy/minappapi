package minappapi

import (
	"encoding/json"
	"testing"

	"github.com/yizenghui/reader"
)

// func Test_MapMake(t *testing.T) {
// 	list, _ := GetList(`http://www.loxow.com/97371.shtml`)
// 	t.Fatal(list.Links[0:5])
// }

func Test_GetPostChapter(t *testing.T) {
	list := GetPostChapter(`http://book.zongheng.com/showchapter/523438.html`)

	b, err := json.Marshal(list)
	if err != nil {
		t.Fatal(err)
	}
	// t.Fatal(string(b))

	var jlist []reader.Link

	json.Unmarshal([]byte(string(b)), &jlist)

	t.Fatal(jlist[0])
}
