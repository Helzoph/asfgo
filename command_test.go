package asfgo

import "testing"

func TestAddLicense(t *testing.T) {
	asf := NewASF("127.0.0.1:1242", "")
	resp, err := asf.AppIDs([]uint32{2999250}).AddLicense()
	if err != nil {
		t.Error(err)
	}

	if !resp.Success {
		t.Error(resp.Message)
	}
}
