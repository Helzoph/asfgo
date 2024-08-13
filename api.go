package asfgo

import (
	"fmt"
	"net/url"
)

func (a *ASF) addLicenseURL() string {
	url := url.URL{
		Scheme: "http",
		Host:   a.host,
		Path:   fmt.Sprintf("/API/Bot/%s/AddLicense", a.botName),
	}

	return url.String()
}

func (a *ASF) commandURL() string {
	url := url.URL{
		Scheme: "http",
		Host:   a.host,
		Path:   "/Api/Command",
	}

	return url.String()
}
