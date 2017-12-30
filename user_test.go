package minappapi

import "testing"

func Test_GetOpenID(t *testing.T) {
	u2, _ := GetOpenID("00397Gwx0b39sj1Pm3wx0guNwx097GwP")
	t.Fatal(u2)
}

func Test_SendPostUpdateMSG(t *testing.T) {
	u2 := SendPostUpdateMSG("oe9Hq0GwS4umXENTCn4lJgxaNVog", "af1a0cdf6adfbf4030358fc2b4264d24", "tttt", "ccccc", "")
	t.Fatal(u2)
}
