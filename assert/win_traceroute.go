//go:build windows

package assert

import (
	"github.com/zarthus/networktest/v2/assert/lib"
)

func TracerouteV4(args lib.Args) lib.Result {
	return lib.ToResult(lib.Warn, "Traceroute using v4 IP", "Traceroute is not supported on windows", nil)
}

func TracerouteV6(args lib.Args) lib.Result {
	return lib.ToResult(lib.Warn, "Traceroute using v6 IP", "Traceroute is not supported on windows", nil)
}
