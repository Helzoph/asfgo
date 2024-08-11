package asfgo

type ASF struct {
	addr    string
	auth    string
	botName string
	appIDs  []uint32
	subIDs  []uint32
}

func NewASF(addr, auth string) *ASF {
	return &ASF{
		addr: addr,
		auth: auth,
	}
}

func (a *ASF) Bot(botName string) *ASF {
	a.botName = botName
	return a
}

func (a *ASF) AppIDs(appIDs []uint32) *ASF {
	a.appIDs = appIDs
	return a
}

func (a *ASF) SubIDs(subIDs []uint32) *ASF {
	a.subIDs = subIDs
	return a
}
