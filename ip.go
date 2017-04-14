// Package ip provides function that checks if an IP is private.
package ip

import (
	"log"
	"net"
)

var privateIPNets [3]*net.IPNet

func init() {
	var err error
	for i, cidr := range []string{"192.168.0.0/16", "10.0.0.0/8", "176.16.0.0/12"} {
		_, privateIPNets[i], err = net.ParseCIDR(cidr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// IsPrivate returns true if the IP is well-formed and if it's part of a
// private network.
func IsPrivate(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	for _, ipnet := range privateIPNets {
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false

}
