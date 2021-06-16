package main

import (
	"encoding/base64"
	"encoding/json"
)

func initialize(Task *MonitorSession) {

	//getting our login token and IV for first time
	getToken(Task)
	generatedIV := generateIV(16)
	encryptedData := AESEncrypt(Task.oauthData.OauthVerificationToken, oauthSecret, generatedIV)
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
	Task.oauthData.OauthVerificationChallenge = generatedIV + encryptedString

	//old dataset so we have something to compare with
	prodResponse := make(chan []byte)
	defer close(prodResponse)
	go getLatest(Task, prodResponse)
	products := <-prodResponse
	json.Unmarshal(products, &Task.oldResponse)

	if !monitor(Task) {
		initialize(Task) //re-init when monitor goes down - loops itself approach
	}
}
