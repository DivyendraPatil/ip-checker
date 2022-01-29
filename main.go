package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var ipRangesApi = map[string]string{
	"GoogleCloud": "www.gstatic.com/ipranges/cloud.json",
	"Google":      "www.gstatic.com/ipranges/goog.json",
	"GoogleBot":   "developers.google.com/search/apis/ipranges/googlebot.json",
	"Fastly":      "api.fastly.com/public-ip-list",
	"Atlassian":   "ip-ranges.atlassian.com",
	"Amazon":      "ip-ranges.amazonaws.com/ip-ranges.json",
}

func main() {

	input := os.Args[1]
	incomingIp := strings.TrimSuffix(input, "\n")

	if net.ParseIP(incomingIp) == nil {
		panic("invalid Ip")
	}

	found := false
	for originName, apiHttp := range ipRangesApi {
		jsonByte := getJsonBody(apiHttp)

		if checkRange(originName, incomingIp, jsonByte) == originName {
			fmt.Printf("This is a %s IP address \n", originName)
			found = true
			break
		}
	}

	if found == false {
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
		if google(body, incomingIp) {
			return originName
		}
	case "Fastly":
		if fastly(body, incomingIp) {
			return originName
		}
	case "Amazon":
		if amazon(body, incomingIp) {
			return originName
		}
	case "Atlassian":
		if atlassian(body, incomingIp) {
			return originName
		}
	}

	return "Not Google"
}

func checkIfIpInRange(cidrRange string, incomingIp string) bool {
	_, subnet, _ := net.ParseCIDR(cidrRange)
	if subnet.Contains(net.ParseIP(incomingIp)) {
		return true
	}
	return false
}
