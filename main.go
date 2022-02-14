package main

import (
	"fmt"
	"ip-checker/api"
	"net"
	"os"
	"strings"
)

var ipRangesApi = map[string]string{
	"GoogleCloud": "https://www.gstatic.com/ipranges/cloud.json",
	"Google":      "https://www.gstatic.com/ipranges/goog.json",
	"GoogleBot":   "https://developers.google.com/search/apis/ipranges/googlebot.json",
	"Fastly":      "https://api.fastly.com/public-ip-list",
	"Atlassian":   "https://ip-ranges.atlassian.com",
	"Amazon":      "https://ip-ranges.amazonaws.com/ip-ranges.json",
	"Bing":        "https://www.bing.com/toolbox/bingbot.json",
}

func main() {

	input := os.Args[1]
	incomingIp := strings.TrimSuffix(input, "\n")

	if net.ParseIP(incomingIp) == nil {
		panic("invalid Ip")
	}

	ipFound := false
	for originName, apiHttp := range ipRangesApi {
		jsonByte := getJsonBody(apiHttp)

		if checkRange(originName, incomingIp, jsonByte) == originName {
			fmt.Printf("This is a %s IP address \n", originName)
			ipFound = true
			break
		}
	}

	if ipFound == false {
		fmt.Println("This ip is not from Google")
	}
}

func checkRange(originName string, incomingIp string, body []byte) string {

	switch originName {
	case "GoogleCloud":
		fallthrough
	case "Google":
		fallthrough
	case "GoogleBot":
		if api.Google(body, incomingIp) {
			return originName
		}
	case "Fastly":
		if api.Fastly(body, incomingIp) {
			return originName
		}
	case "Amazon":
		if api.Amazon(body, incomingIp) {
			return originName
		}
	case "Atlassian":
		if api.Atlassian(body, incomingIp) {
			return originName
		}
	case "Bing":
		if api.Bing(body, incomingIp) {
			return originName
		}
	}

	return "Not Google"
}
