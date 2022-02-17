package api

import (
	"encoding/json"
	"fmt"
)

type CloudflareIpRangesStruct struct {
	Result struct {
		Ipv4Cidrs []string `json:"ipv4_cidrs"`
		Ipv6Cidrs []string `json:"ipv6_cidrs"`
		Etag      string   `json:"etag"`
	} `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

func Cloudflare(body []byte, incomingIp string) bool {
	var result CloudflareIpRangesStruct
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON from api response")
	}

	for _, rec := range result.Result.Ipv4Cidrs {
		if checkIfIpInRange(rec, incomingIp) == true {
			return true
		}
	}
	for _, rec := range result.Result.Ipv6Cidrs {
		if checkIfIpInRange(rec, incomingIp) == true {
			return true
		}
	}
	return false
}
