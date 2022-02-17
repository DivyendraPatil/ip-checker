package api

import (
	"encoding/json"
	"fmt"
)

type GoogleIpRangesStruct struct {
	SyncToken    string `json:"syncToken,omitempty"`
	CreationTime string `json:"creationTime"`
	Prefixes     []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
		Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
		Service    string `json:"service,omitempty"`
		Scope      string `json:"scope,omitempty"`
	} `json:"prefixes"`
}

func Google(body []byte, incomingIp string) (string, bool) {

	var result GoogleIpRangesStruct
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON from api response")
	}
	cidrRange := result.Prefixes

	for _, rec := range cidrRange {
		if rec.Ipv4Prefix != "" {
			if checkIfIpInRange(rec.Ipv4Prefix, incomingIp) == true {
				return metadata(rec.Scope), true
			}
		} else if rec.Ipv6Prefix != "" {
			if checkIfIpInRange(rec.Ipv6Prefix, incomingIp) == true {
				return metadata(rec.Scope), true
			}
		}
	}

	return "", false
}
