package assert

import (
	"github.com/zarthus/networktest/v2/assert/lib"
	"net"
)

func DnsLookupIp(args lib.Args) lib.Result {
	host, _ := lib.SplitHostAndPort(lib.GetHost(args, "w3.org:443"))

	resolved, err := net.LookupIP(host)

	var resolves []string

	if err == nil {
		for _, resolve := range resolved {
			resolves = append(resolves, resolve.String())
		}
	}

	return lib.ResultByError("DNS Lookup (A / AAAA)", "Resolve DNS of "+host, err, resolves)
}

func DnsLookupMx(args lib.Args) lib.Result {
	host, _ := lib.SplitHostAndPort(lib.GetHost(args, "w3.org:443"))

	resolved, err := net.LookupMX(host)

	var resolves []string

	if err == nil {
		for _, resolve := range resolved {
			resolves = append(resolves, resolve.Host)
		}
	}

	return lib.ResultByError("DNS Lookup (MX)", "Resolve DNS of "+host, err, resolves)
}
