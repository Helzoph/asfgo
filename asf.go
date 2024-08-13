package asfgo

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type ASF struct {
	host    string
	auth    string
	botName string
	appIDs  []string
	subIDs  []string
}

func NewASF(addr, auth string) *ASF {
	return &ASF{
		host: addr,
		auth: auth,
	}
}

func (a *ASF) CheckConnect() (err error) {
	url := url.URL{
		Scheme: "http",
		Host:   a.host,
		Path:   "/Api/ASF",
	}

	resp, err := a.get(url.String())
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to connect: %s", resp.Message)
	}

	return
}

func (a *ASF) Bot(botName string) *ASF {
	a.botName = botName
	return a
}

func (a *ASF) AppIDs(appIDs any) *ASF {
	tp := reflect.TypeOf(appIDs)
	if tp.Kind() == reflect.Slice {
		switch tp.Elem().Kind() {
		case reflect.String:
			a.AppIDsStr(appIDs.([]string))
		case reflect.Int:
			a.AppIDsInt(appIDs.([]int))
		default:
			panic(ErrInvalidType)
		}
		return a
	}

	switch tp.Kind() {
	case reflect.String:
		a.AppIDsStr([]string{appIDs.(string)})
	case reflect.Int:
		a.AppIDsInt([]int{appIDs.(int)})
	default:
		panic(ErrInvalidType)
	}

	return a
}

func (a *ASF) AppIDsStr(appIDs []string) *ASF {
	a.appIDs = appIDs
	return a
}

func (a *ASF) AppIDsInt(appIDsInt []int) *ASF {
	appIDs := make([]string, len(appIDsInt))
	for i, v := range appIDsInt {
		appIDs[i] = strconv.Itoa(v)
	}
	a.appIDs = appIDs
	return a
}

func (a *ASF) SubIDs(subIDs any) *ASF {
	tp := reflect.TypeOf(subIDs)
	if tp.Kind() == reflect.Slice {
		switch tp.Elem().Kind() {
		case reflect.String:
			a.SubIDsStr(subIDs.([]string))
		case reflect.Int:
			a.SubIDsInt(subIDs.([]int))
		default:
			panic(ErrInvalidType)
		}
		return a
	}

	switch tp.Kind() {
	case reflect.String:
		a.SubIDsStr([]string{subIDs.(string)})
	case reflect.Int:
		a.SubIDsInt([]int{subIDs.(int)})
	default:
		panic(ErrInvalidType)
	}

	return a
}

func (a *ASF) SubIDsStr(subIDs []string) *ASF {
	a.subIDs = subIDs
	return a
}

func (a *ASF) SubIDsInt(subIDsInt []int) *ASF {
	subIDs := make([]string, len(subIDsInt))
	for i, v := range subIDsInt {
		subIDs[i] = strconv.Itoa(v)
	}
	a.subIDs = subIDs
	return a
}
