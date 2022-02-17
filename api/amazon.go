package api

import (
	"encoding/json"
	"fmt"
)

type AmazonIpRangesStruct struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix           string `json:"ip_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"prefixes"`
	Ipv6Prefixes []struct {
		Ipv6Prefix         string `json:"ipv6_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"ipv6_prefixes"`
}

func Amazon(body []byte, incomingIp string) (string, bool) {
	var result AmazonIpRangesStruct
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON from api response")
	}

	for _, rec := range result.Prefixes {
		if checkIfIpInRange(rec.IPPrefix, incomingIp) == true {
			return metadata(rec.Region), true
		}
	}
	for _, rec := range result.Ipv6Prefixes {
		if checkIfIpInRange(rec.Ipv6Prefix, incomingIp) == true {
			return metadata(rec.Region), true
		}
	}
	return "", false
}
