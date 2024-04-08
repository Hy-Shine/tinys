package netx

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidIP(t *testing.T) {
	cases := []struct {
		in  string
		exp bool
	}{
		{in: "::1", exp: true},
		{in: "127.0.0.1", exp: true},
		{in: "192.168..", exp: false},
		{in: "192.168..1", exp: false},
	}

	for _, v := range cases {
		assert.Equal(t, v.exp, IsIP(v.in))
	}
}

func TestIsValidIPv4(t *testing.T) {
	cases := []struct {
		in  string
		exp bool
	}{
		{in: "::1", exp: false},
		{in: "127.0.0.1", exp: true},
		{in: "192.168.0.", exp: false},
		{in: "192.168..1", exp: false},
		{in: "", exp: false},
		{in: "127.0.0.1.", exp: false},
	}
	for _, v := range cases {
		assert.Equal(t, v.exp, IsIPv4(v.in), fmt.Sprintf("case: %s not pass", v.in))
	}
}

func TestIsValidIPv6(t *testing.T) {
	cases := []struct {
		in  string
		exp bool
	}{
		{in: "::1", exp: true},
		{in: "127.0.0.1", exp: false},
		{in: "1080::8:800:200C:417A", exp: true},
		{in: "", exp: false},
		{in: "FF01::101", exp: true},
		{in: "1080:0:0:0:8:800:200C:417A", exp: true},
	}
	for _, v := range cases {
		assert.Equal(t, v.exp, IsIPv6(v.in), fmt.Sprintf("case: %s NOT PASS", v.in))
	}
}

func TestCIDRContainsIP(t *testing.T) {
	cases := []struct {
		cidr, ip string
		exp      bool
	}{
		{cidr: "192.168.1.0/24", ip: "192.168.1.100", exp: true},
		{cidr: "192.168.1.0/24", ip: "192.168.2.100", exp: false},
		{cidr: "172.17.0.0/16", ip: "172.17.0.10", exp: true},
		{cidr: "10.0.0.0/24", ip: "127.0.0.1", exp: false},
	}
	for _, v := range cases {
		assert.Equal(t, v.exp, CIDRContainsIP(v.cidr, v.ip), fmt.Sprintf("cidr:%s ip:%s NOT PASS", v.cidr, v.ip))
	}
}

func TestToLong(t *testing.T) {
	ip := net.ParseIP("192.168.0.1")
	fmt.Println(IPv4ToLong(ip))
}

func TestLongToIP(t *testing.T) {
	longIP := IPv4ToLong(net.ParseIP("192.168.0.1"))
	fmt.Println(LongToIP(longIP))
}

func BenchmarkToLongOr(t *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	for i := 0; i < t.N; i++ {
		IPv4ToLong(ip)
	}
}
