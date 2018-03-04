package gogetmyip

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const ipUrl = "http://jsonip.com"

func getMapBody(bodyByte []byte) (map[string]interface{}, error) {
	var data map[string]interface{}

	if err := json.Unmarshal(bodyByte, &data); err != nil {
		return make(map[string]interface{}), err
	}

	return data, nil
}

func grabHttp(url string) ([]byte, error) {

	ipClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return make([]byte, 0, 0), err
	}

	res, getErr := ipClient.Do(req)
	if getErr != nil {
		return make([]byte, 0, 0), getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return make([]byte, 0, 0), readErr
	}

	return body, nil
}

func getIpKey(jsonip map[string]interface{}) (string, error) {
	ip, exists := jsonip["ip"]
	// todo: validate ip format
	if exists {
		value, isString := ip.(string)
		if isString {
			return value, nil
		} else {
			return "", errors.New("Json value is not a string")
		}

	} else {
		return "", errors.New("Invalid json format")
	}
}

func GetIp() (string, error) {
	body, reqErr := grabHttp(ipUrl)
	if reqErr != nil {
		return "", reqErr
	}
	ipmap, parseErr := getMapBody(body)
	if reqErr != nil {
		return "", parseErr
	}
	ip, err := getIpKey(ipmap)
	return ip, err
}
