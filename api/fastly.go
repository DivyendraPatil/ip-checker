package api

import (
	"encoding/json"
	"fmt"
)

type FastlyIpRangesStruct struct {
	Addresses     []string `json:"addresses"`
	Ipv6Addresses []string `json:"ipv6_addresses"`
}

func Fastly(body []byte, incomingIp string) bool {
	var result FastlyIpRangesStruct
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON from api response")
	}

	for _, rec := range result.Addresses {
		if checkIfIpInRange(rec, incomingIp) == true {
			return true
		}
	}
	for _, rec := range result.Ipv6Addresses {
		if checkIfIpInRange(rec, incomingIp) == true {
			return true
		}
	}
	return false
}
