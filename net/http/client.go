package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PostJson(url string, parms interface{}, ret interface{}) error {
	client := &http.Client{}
	data, err := json.Marshal(parms)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return err
	}
	return nil
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
