package api

import (
	"net"
	"strings"
)

func checkIfIpInRange(cidrRange string, incomingIp string) bool {

	_, subnet, err := net.ParseCIDR(strings.TrimSpace(cidrRange))
	if subnet.Contains(net.ParseIP(incomingIp)) && err != nil {
		return true
	}
	return false
}
