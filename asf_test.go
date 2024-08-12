package asfgo

import "testing"

func TestAppIDs(t *testing.T) {
	asf := NewASF("", "")
	asf.AppIDs(0.1)
}
