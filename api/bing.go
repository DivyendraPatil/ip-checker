package api

import (
	"encoding/json"
	"fmt"
)

type BingIpRangesStruct struct {
	CreationTime string `json:"creationTime"`
	Prefixes     []struct {
		Ipv4Prefix string `json:"ipv4Prefix"`
	} `json:"prefixes"`
}

func Bing(body []byte, incomingIp string) bool {
	var result BingIpRangesStruct
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON from api response")
	}

	for _, rec := range result.Prefixes {
		if checkIfIpInRange(rec.Ipv4Prefix, incomingIp) == true {
			return true
		}
	}
	return false
}
