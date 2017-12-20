package minappapi

import "testing"

func Test_GetOpenID(t *testing.T) {
	u2, _ := GetOpenID("00397Gwx0b39sj1Pm3wx0guNwx097GwP")
	t.Fatal(u2)
}
