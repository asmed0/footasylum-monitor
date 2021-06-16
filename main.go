package main

import (
	http "github.com/useflyent/fhttp"
)

func main() {
	Task := &MonitorSession{
		oauthData:   &TokenResponse{},
		newResponse: &ProductsResponse{},
		oldResponse: &ProductsResponse{},
		Client:      &http.Client{},
	}

	initialize(Task)
}
