package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type AzureIpRangesStruct []struct {
	ServiceTagID    string `json:"serviceTagId,omitempty"`
	IPAddress       string `json:"ipAddress,omitempty"`
	IPAddressPrefix string `json:"ipAddressPrefix,omitempty"`
	Region          string `json:"region,omitempty"`
	SystemService   string `json:"systemService,omitempty"`
	NetworkFeatures string `json:"networkFeatures,omitempty"`
}

var azureCustomApi = "https://www.azurespeed.com/api/ipinfo?ipAddressOrUrl="

func Azure(incomingIp string) (string, bool) {
	resp, err := http.Get(azureCustomApi + incomingIp)
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

	var result AzureIpRangesStruct
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON from api response")
	}

	for _, rec := range result {
		if checkIfIpInRange(rec.IPAddressPrefix, incomingIp) == true {
			return rec.ServiceTagID + " | " + rec.SystemService, true
		}
	}

	return "", false
}
