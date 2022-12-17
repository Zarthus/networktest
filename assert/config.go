package assert

import (
	"github.com/zarthus/networktest/v2/assert/lib"
	"net"
	"strings"
)

func CheckIpVersions(args lib.Args) lib.Result {
	hasIpv4, hasIpv6 := false, false
	var guidance []string

	ifaces, err := net.InterfaceAddrs()
	if err != nil {
		return lib.ResultByError("Assigned IP(s)", "IPv4 and IPv6 assigned", err, nil)
	}

	var ip4s []string
	var ip6s []string

	for _, iface := range ifaces {
		split := strings.Split(iface.String(), "/")

		ip := net.ParseIP(split[0])

		if ip == nil || lib.IsLocal(ip) {
			continue
		}

		if strings.Contains(ip.String(), ".") {
			hasIpv4 = true
			ip4s = append(ip4s, ip.String())
		} else if strings.Contains(ip.String(), ":") {
			hasIpv6 = true
			ip6s = append(ip6s, ip.String())
		}
	}

	stat := lib.Ok
	if !hasIpv4 && !hasIpv6 {
		guidance = append(guidance, "Could not find public IP addresses")
		stat = lib.Warn
	} else if !hasIpv4 {
		stat = lib.Warn
		guidance = append(guidance, "Missing IPv4 connectivity")
		guidance = append(guidance, ip6s...)
	} else if !hasIpv6 {
		stat = lib.Warn
		guidance = append(guidance, "Missing IPv6 connectivity")
		guidance = append(guidance, ip4s...)
	} else {
		guidance = append(guidance, ip4s...)
		guidance = append(guidance, ip6s...)
	}

	return lib.ToResult(stat, "Assigned IP(s)", "Check local IP configuration", guidance)
}

func CheckInterfaces(args lib.Args) lib.Result {
	ifaces, err := net.Interfaces()
	var guidance []string

	if err == nil {
		for _, iface := range ifaces {
			addrs, _ := iface.Addrs()

			if len(addrs) > 0 {
				var addrstrs []string
				for _, addr := range addrs {
					addrstrs = append(addrstrs, addr.String())
				}
				guidance = append(guidance, iface.Name+" ["+strings.Join(addrstrs, ", ")+"]")
			} else {
				guidance = append(guidance, iface.Name+" (no addresses)")
			}
		}
	}

	return lib.ResultByError("Network Configation", "Looking up interfaces", err, guidance)
}

func CheckPrivilegedUser(args lib.Args) lib.Result {
	rooted := lib.IsRoot()

	if !rooted {
		return lib.ToResult(lib.Warn, "rooted", "Not running as root", []string{"Some lookups may not complete successfully due to requiring a Privileged User"})
	}

	return lib.ToResult(lib.Ok, "Rooted", "Running as root", nil)
}
