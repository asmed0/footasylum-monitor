package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	http "github.com/useflyent/fhttp"
)

func getLatest(Task *MonitorSession, response chan []byte) {
	url := "https://api-e3.nuqlium.com/api/2.0/experience"
	method := "POST"
	payload := strings.NewReader(`pagetype=category&pagekey=new_mens&dataindex=100029&mode=app&settings.layout=0&output=json&settings[page]=1`)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Host", "api-e3.nuqlium.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "Bearer "+Task.oauthData.AccessToken)
	req.Header.Add("oauth-verification-challenge", Task.oauthData.OauthVerificationChallenge)
	req.Header.Add("user-agent", "Footasylum/1355 CFNetwork/1209 Darwin/20.2.0")
	req.Header.Add("accept-language", "sv-se")

	res, err := Task.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode == 200 {
		response <- body
	} else if res.StatusCode == 403 {
		fmt.Println(string(body))
		log.Println("[MONITOR] Access denied", res.Status)
		response <- []byte("err")
	} else if res.StatusCode == 429 {
		log.Println("[MONITOR ] Rate limited", res.Status)
		response <- []byte("err")
	} else {
		log.Println("[MONITOR] Unknown error")
		fmt.Println(string(body))
		response <- []byte("err")
	}
}
