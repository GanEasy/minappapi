// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package minappapi

import (
	"testing"
)

func Test_GetSubscribeCheckList(t *testing.T) {

	var subscribe Subscribe
	subscribes := subscribe.GetSubscribeCheckList()
	t.Fatal(subscribes)
}

func Test_GetSubscribePost(t *testing.T) {

	var post Post
	posts := post.GetSubscribePost()
	t.Fatal(posts)
}
