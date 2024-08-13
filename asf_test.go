package asfgo

import "testing"

func TestAppIDs(t *testing.T) {
	asf := NewASF("", "")
	asf.AppIDs(1)
	asf.AppIDs("1")
	asf.AppIDs([]int{1, 2, 3})
	asf.AppIDs([]string{"1", "2", "3"})
}

func TestCheckConnect(t *testing.T) {
	asf := NewASF("127.0.0.1:1242", "")

	err := asf.CheckConnect()
	if err != nil {
		t.Error(err)
	}
}
