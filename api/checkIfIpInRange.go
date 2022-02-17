package api

import (
	"net"
)

func checkIfIpInRange(cidrRange string, incomingIp string) bool {

	_, subnet, err := net.ParseCIDR(cidrRange)
	if subnet.Contains(net.ParseIP(incomingIp)) && err != nil {
		return true
	}
	return false
}
