package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	http "github.com/useflyent/fhttp"

	"strings"
)

func getToken(Task *MonitorSession) {

	url := "https://api-e3.nuqlium.com/api/token"
	method := "POST"

	payload := strings.NewReader(`username=footasylum-b12cf60a-0db8-40a1-aaf6-04e6ffe4e342&password=553c3198-652a-4355-97b9-d6b0e5a6e15d&grant_type=password`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "api-e3.nuqlium.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Footasylum/1355 CFNetwork/1209 Darwin/20.2.0")
	req.Header.Add("accept-language", "sv-se")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(body, &Task.oauthData)
}
