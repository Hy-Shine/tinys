package netx

import (
	"net"
	"strings"
)

func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsIPv4(ip string) bool {
	if index := strings.IndexByte(ip, '.'); index > 0 {
		return IsIP(ip)
	}
	return false
}

func IsIPv6(ip string) bool {
	if index := strings.IndexByte(ip, ':'); index >= 0 {
		return IsIP(ip)
	}
	return false
}

func CIDRContainsIP(cidr, ip string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	parsedIP := net.ParseIP(ip)
	return ipnet.Contains(parsedIP)
}
