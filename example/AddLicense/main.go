package main

import "github.com/helzoph/asfgo"

func main() {
	asf := asfgo.NewASF("127.0.0.1:1242", "")
	resp, err := asf.AppIDs([]uint32{230410}).AddLicense()
	if err != nil {
		panic(err)
	}

	if !resp.Success {
		panic(resp.Message)
	}
}
