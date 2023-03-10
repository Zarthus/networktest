package main

import (
	"github.com/zarthus/networktest/v2/assert"
	"github.com/zarthus/networktest/v2/assert/lib"
	"os"
	"time"
)

type method string

const (
	dns  method = "dns"
	v4   method = "v4"
	v6   method = "v6"
	misc method = "misc"
)

type item struct {
	method   method
	Runnable func(args lib.Args) lib.Result
}

func runSuite(args lib.Args, parallel bool, colourize bool) {
	suite := createSuite(args)

	if len(suite) == 0 {
		println("error: no tests to run!")
		os.Exit(1)
	}

	if parallel {
		pending := len(suite)
		for _, it := range suite {
			go func(it item) {
				doRun(it, args, colourize)
				pending -= 1
			}(it)
		}

		for pending != 0 {
			time.Sleep(1000)
		}
	} else {
		for _, it := range suite {
			doRun(it, args, colourize)
		}
	}
}

func doRun(it item, args lib.Args, colourize bool) {
	result := it.Runnable(args)

	Write(result, colourize)
}

func createSuite(args lib.Args) []item {
	suite := []item{
		//{method: misc, Runnable: assert.CheckPrivilegedUser},
		{method: misc, Runnable: assert.CheckIpVersions},
		//{method: misc, Runnable: assert.CheckInterfaces},
		{method: dns, Runnable: assert.DnsLookupIp},
		{method: dns, Runnable: assert.DnsLookupMx},
		{method: dns, Runnable: assert.ProbeDns},
		{method: v4, Runnable: assert.ProbeV4},
		{method: v6, Runnable: assert.ProbeV6},
		{method: dns, Runnable: assert.QueryHttpsDns},
		{method: v4, Runnable: assert.QueryHttpV4},
		{method: v6, Runnable: assert.QueryHttpV6},
	}
	suite = append(suite, suiteExtras(args)...)

	return applyFilter(args, suite)
}

func applyFilter(args lib.Args, suite []item) []item {
	if args.V4 && args.V6 {
		return suite
	}

	var filteredSuite []item
	for _, it := range suite {
		if args.V4 && it.method == v4 {
			filteredSuite = append(filteredSuite, it)
		}
		if args.V6 && it.method == v6 {
			filteredSuite = append(filteredSuite, it)
		}
	}

	return filteredSuite
}
