package main

import (
	"fmt"
	"ip-checker/api"
	"net"
	"os"
	"strings"
)

var ipRangesApi = map[string]string{
	"Google Cloud":  "https://www.gstatic.com/ipranges/cloud.json",
	"Google":        "https://www.gstatic.com/ipranges/goog.json",
	"Google Bot":    "https://developers.google.com/search/apis/ipranges/googlebot.json",
	"Fastly":        "https://api.fastly.com/public-ip-list",
	"Atlassian":     "https://ip-ranges.atlassian.com",
	"Amazon":        "https://ip-ranges.amazonaws.com/ip-ranges.json",
	"BingBot":       "https://www.bing.com/toolbox/bingbot.json",
	"Cloudflare":    "https://api.cloudflare.com/client/v4/ips",
	"Akamai":        "https://raw.githubusercontent.com/SecOps-Institute/Akamai-ASN-and-IPs-List/master/akamai_ip_cidr_blocks.lst",
	"Digital Ocean": "https://raw.githubusercontent.com/SecOps-Institute/Digitalocean-ASN-and-IPs-List/master/digitalocean_ip_cidr_blocks.lst",
	"LinkedIn":      "https://raw.githubusercontent.com/SecOps-Institute/LinkedInIPLists/master/linkedin_ipv4_cidr_blocks.lst",
	"Twitter":       "https://raw.githubusercontent.com/SecOps-Institute/TwitterIPLists/master/twitter_ipv4_cidr_blocks.lst",
	"Facebook":      "https://raw.githubusercontent.com/SecOps-Institute/FacebookIPLists/master/facebook_ipv4_cidr_blocks.lst",
}

func main() {

	input := os.Args[1]
	incomingIp := strings.TrimSuffix(input, "\n")

	if net.ParseIP(incomingIp) == nil {
		panic("invalid Ip")
	}

	ipFound := false
	for originName, apiHttp := range ipRangesApi {
		response := getResponseBody(apiHttp)

		fmt.Println("\n", originName, "\n")
		origin, check := checkRange(originName, incomingIp, response)
		if check == true {
			fmt.Println(origin)
			ipFound = true
		}
	}

	// Because Microsoft does not have a cidr api
	scope, valid := api.Azure(incomingIp)
	if valid {
		fmt.Println("Microsoft Azure | " + scope)
	}

	if ipFound == false {
		fmt.Println("Unknown IP Address")
	}
}

func checkRange(originName string, incomingIp string, body []byte) (string, bool) {

	switch originName {
	case "Google Cloud":
		fallthrough
	case "Google":
		fallthrough
	case "Google Bot":
		scope, valid := api.Google(body, incomingIp)
		if valid {
			return originName + " | " + scope, true
		}
	case "Fastly":
		if api.Fastly(body, incomingIp) {
			return originName, true
		}
	case "Amazon":
		scope, valid := api.Amazon(body, incomingIp)
		if valid {
			return originName + " | " + scope, true
		}
	case "Atlassian":
		if api.Atlassian(body, incomingIp) {
			return originName, true
		}
	case "Bing":
		if api.Bing(body, incomingIp) {
			return originName, true
		}
	case "Cloudflare":
		if api.Cloudflare(body, incomingIp) {
			return originName, true
		}
	case "Akamai":
		fallthrough
	case "Digital Ocean":
		fallthrough
	case "LinkedIn":
		fallthrough
	case "Twitter":
		fallthrough
	case "Spamhaus":
		fallthrough
	case "Facebook":
		if api.Standard(body, incomingIp) {
			return originName, true
		}
	}

	return "Unknown Address", false
}
