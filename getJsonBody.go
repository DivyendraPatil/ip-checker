package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func getJsonBody(apiHttp string) []byte {

	sslURL := strings.Split(apiHttp, ".com")

	// Validate SSL for api endpoint
	connection, err := tls.Dial("tcp", sslURL[0]+".com:443", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: \n" + err.Error())
	}

	err = connection.VerifyHostname(sslURL[0] + ".com")
	if err != nil {
		panic("Hostname doesn't match with certificate: " + err.Error())
	}

	resp, err := http.Get("https://" + apiHttp)
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
