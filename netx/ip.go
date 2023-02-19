package netx

import (
	"net"
	"strings"
)

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsValidIPv4(ip string) bool {
	if !strings.Contains(ip, ".") || strings.Contains(ip, ":") {
		return false
	}
	return IsValidIP(ip)
}

func IsValidIPv6(ip string) bool {
	if !strings.Contains(ip, ":") || strings.Contains(ip, ".") {
		return false
	}
	return IsValidIP(ip)
}

func CIDRContainsIP(cidr, ip string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	parsedIP := net.ParseIP(ip)
	return ipnet.Contains(parsedIP)
}
