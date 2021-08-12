package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostJsonWithHeaders(url string, parms interface{},
	headers map[string]string, ret interface{}) (*http.Request, error) {
	client := &http.Client{}
	data, err := json.Marshal(parms)
	if err != nil {
		return nil, fmt.Errorf("failed Marshal, err:%v, parms:%+v", err, parms)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed NewRequest, err:%v, data:%+v", err, string(data))
	}

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "name=anny")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return req, fmt.Errorf("failed NewRequest, err:%v, data:%+v", err, string(data))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return req, fmt.Errorf("failed ReadAll, err:%v", err)
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return req, fmt.Errorf("failed Unmarshal, err:%v, body:%v", err, string(body))
	}
	return req, nil
}

func PostJson(url string, parms interface{}, ret interface{}) (*http.Request, error) {
	return PostJsonWithHeaders(url, parms, nil, ret)
}

func PostJsonReturnMap(url string, parms interface{}) (interface{}, error) {
	client := &http.Client{}
	data, err := json.Marshal(parms)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p interface{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
