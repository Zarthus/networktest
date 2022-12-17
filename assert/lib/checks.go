package lib

import (
	"net"
	"net/netip"
	"os/user"
	"regexp"
	"strconv"
)

func IsRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		return false
	}
	return currentUser.Username == "root"
}

func IsLocal(ip net.IP) bool {
	prefixes := []netip.Prefix{
		netip.MustParsePrefix("127.0.0.1/8"),
		netip.MustParsePrefix("10.0.0.0/8"),
		netip.MustParsePrefix("172.16.0.0/12"),
		netip.MustParsePrefix("192.168.0.0/16"),
		netip.MustParsePrefix("fd00::/8"),
		netip.MustParsePrefix("fe80::/8"),
		netip.MustParsePrefix("::1/64"),
	}

	addr := netip.MustParseAddr(ip.String())
	for _, network := range prefixes {
		if network.Contains(addr) {
			return true
		}
	}
	return false
}

func SplitHostAndPort(hostname string) (string, int) {
	host := hostname
	port := 443

	re, err := regexp.Compile(`(.*]?):(\d+)$`)

	if err != nil {
		return host, port
	}

	result := re.FindStringSubmatch(hostname)
	if len(result) > 1 {
		host = result[1]
	}
	if len(result) > 2 {
		port, err = strconv.Atoi(result[2])
		if err != nil {
			port = 443
		}
	} else {
		port = 443
	}

	return host, port
}

func GetHost(args Args, fallback string) string {
	if args.Host != "" && args.Host != "auto" {
		return args.Host
	}
	return fallback
}
