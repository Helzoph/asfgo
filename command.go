package asfgo

func (a *ASF) AddLicense() (resp *response, err error) {
	data := struct {
		Apps     []uint32 `json:"Apps"`
		Packages []uint32 `json:"Packages"`
	}{
		Apps:     a.appIDs,
		Packages: a.subIDs,
	}

	if a.botName == "" {
		a.botName = "asf"
	}

	resp, err = a.post(a.addLicenseURL(), data)
	if err != nil {
		return
	}

	return
}

func (a *ASF) Command(cmd string) (resp *response, err error) {
	data := struct {
		Command string `json:"Command"`
	}{
		Command: cmd,
	}

	resp, err = a.post(a.commandURL(), data)
	if err != nil {
		return
	}

	return
}
