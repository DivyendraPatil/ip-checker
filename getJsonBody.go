package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func getJsonBody(apiHttp string) []byte {

	resp, err := http.Get(apiHttp)
	if err != nil {
		fmt.Println("No response from request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	return body
}
