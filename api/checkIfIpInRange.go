package api

import "net"

func checkIfIpInRange(cidrRange string, incomingIp string) bool {
	_, subnet, _ := net.ParseCIDR(cidrRange)
	if subnet.Contains(net.ParseIP(incomingIp)) {
		return true
	}
	return false
}
