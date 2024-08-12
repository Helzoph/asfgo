# ASFGo

Controlling ASF by calling the ASF API

## Usage

Make sure ASF is enabled before use [IPC](https://github.com/JustArchiNET/ArchiSteamFarm/wiki/IPC)

```bash
go get github.com/helzoph/asfgo
```

### AddLicense

```go
package main

import "github.com/helzoph/asfgo"

func main() {
	asf := asfgo.NewASF("127.0.0.1:1242", "")
	resp, err := asf.AppIDs([]string{230410}).AddLicense()
	if err != nil {
		panic(err)
	}

	if !resp.Success {
		panic(resp.Message)
	}
}
```

## TODO

- [x] AddLicense
- [x] Command
- [ ] [ASFEnhance](https://github.com/chr233/ASFEnhance)