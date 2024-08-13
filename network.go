package asfgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type Response struct {
	Message string      `json:"Message"`
	Success bool        `json:"Success"`
	Result  interface{} `json:"Result"`
}

func (a *ASF) post(url string, data interface{}) (result *Response, err error) {
	tp := reflect.TypeOf(data)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}

	if tp.Kind() != reflect.Struct {
		return &Response{}, fmt.Errorf("data must be a struct")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", a.auth)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (a *ASF) get(url string) (result *Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", a.auth)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}
