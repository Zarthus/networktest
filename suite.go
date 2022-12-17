//go:build !windows

package main

import (
	"github.com/zarthus/networktest/v2/assert"
	"github.com/zarthus/networktest/v2/assert/lib"
)

func createSuite(args lib.Args) []item {
	return applyFilter(args, []item{
		//{method: misc, Runnable: assert.CheckPrivilegedUser},
		{method: misc, Runnable: assert.CheckIpVersions},
		//{method: misc, Runnable: assert.CheckInterfaces},
		{method: dns, Runnable: assert.DnsLookupIp},
		{method: dns, Runnable: assert.DnsLookupMx},
		{method: dns, Runnable: assert.ProbeDns},
		{method: v4, Runnable: assert.ProbeV4},
		{method: v6, Runnable: assert.ProbeV6},
		{method: v4, Runnable: assert.TracerouteV4},
		{method: v6, Runnable: assert.TracerouteV6},
	})
}
