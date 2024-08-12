# ASFGo

[English](./README_en.md)

通过调用 ASF API 来控制 ASF

## Usage

使用之前请确保 ASF 已开启 [IPC](https://github.com/JustArchiNET/ArchiSteamFarm/wiki/IPC)

```bash
go get github.com/helzoph/asfgo
```

## Example

### AddLicense

```go
package main

import "github.com/helzoph/asfgo"

func main() {
	asf := asfgo.NewASF("127.0.0.1:1242", "")
	resp, err := asf.AppIDs([]string{"230410"}).AddLicense()
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