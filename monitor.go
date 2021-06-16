package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

func monitor(Task *MonitorSession) bool {

	//making new IV for each time
	generatedIV := generateIV(16)
	encryptedData := AESEncrypt(Task.oauthData.OauthVerificationToken, oauthSecret, generatedIV)
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
	Task.oauthData.OauthVerificationChallenge = generatedIV + encryptedString
	//new dataset
	response := make(chan []byte)
	defer close(response)
	go getLatest(Task, response)
	products := <-response
	if string(products) != "err" { //should only happen when bearer token expires, we renew
		json.Unmarshal(products, &Task.newResponse)
		if Task.newResponse.Data.Products.Products[0].Product.Productid ==
			Task.oldResponse.Data.Products.Products[0].Product.Productid {
			fmt.Println("No new products loaded")
			time.Sleep(3000 * time.Millisecond)
			monitor(Task)
		} else {
			fmt.Printf("New product loaded with pid: %s\n", Task.newResponse.Data.Products.Products[0].Product.Productid)
		}
	}
	return false
}
