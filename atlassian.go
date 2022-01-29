package main

import (
	"encoding/json"
	"fmt"
)

type AtlassianIpRangesStruct struct {
	CreationDate string `json:"creationDate"`
	SyncToken    int    `json:"syncToken"`
	Items        []struct {
		Network   string   `json:"network"`
		MaskLen   int      `json:"mask_len"`
		Cidr      string   `json:"cidr"`
		Mask      string   `json:"mask"`
		Region    []string `json:"region"`
		Product   []string `json:"product"`
		Direction []string `json:"direction"`
	} `json:"items"`
}

func atlassian(body []byte, incomingIp string) bool {
	var result AtlassianIpRangesStruct
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON from api response")
	}

	for _, rec := range result.Items {
		if checkIfIpInRange(rec.Cidr, incomingIp) == true {
			return true
		}
	}
	return false
}
