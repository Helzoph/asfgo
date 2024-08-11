package asfgo

import "fmt"

func (a *ASF) addLicenseURL() string {
	return fmt.Sprintf("http://%s/Api/Bot/%s/AddLicense", a.addr, a.botName)
}

func (a *ASF) commandURL() string {
	return fmt.Sprintf("http://%s/Api/Command", a.addr)
}
