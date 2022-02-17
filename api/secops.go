package api

import (
	"bufio"
	"strings"
)

func Standard(body []byte, incomingIp string) bool {
	cidrs := string(body)

	scanner := bufio.NewScanner(strings.NewReader(cidrs))
	for scanner.Scan() {
		cidr := scanner.Text()
		if checkIfIpInRange(cidr, incomingIp) == true {
			return true
		}
	}

	return false
}
