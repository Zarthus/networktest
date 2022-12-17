package main

import (
	"flag"
	"github.com/zarthus/networktest/v2/assert/lib"
	"os"
)

func main() {
	only := flag.Int("only", 0, "Only test IPv4 or IPv6 connectivity")
	timeout := flag.Int("timeout", 3, "TimeoutSeconds in seconds")
	targetHost := flag.String("target", "auto", "Which host to test against")
	dontColourize := flag.Bool("no-col", false, "Do not output colours")
	parallel := flag.Bool("parallel", false, "Run commands in parallel")
	hlp := flag.Bool("help", false, "Help text")
	ver := flag.Bool("version", false, "Version text")
	flag.Parse()

	if *hlp {
		help()
		os.Exit(0)
	}

	if *ver {
		version()
		os.Exit(0)
	}

	args := validate(only, timeout, targetHost)

	version()
	runSuite(args, *parallel, !*dontColourize)
}

func validate(only *int, timeout *int, targetHost *string) lib.Args {
	if *timeout < 1 {
		println("error: TimeoutSeconds is too low (under 1 second)")
		help()
		os.Exit(2)
	}

	return lib.Args{
		V4:             only == nil || *only == 0 || *only == 4,
		V6:             only == nil || *only == 0 || *only == 6,
		TimeoutSeconds: *timeout,
		Host:           *targetHost,
	}
}

const versionStr = "0.1"

func help() {
	println(
		"network tester v" + versionStr + " - tests network connectivity\n\n" +
			"  --only=4|6       only test ipv4 or ipv6 connectivity [default both]\n" +
			"  --timeout=int    timeout in seconds [default 3]\n" +
			"  --target=string  runs some testcases against a specific hostname [default \"auto\"]\n" +
			"                     please be aware that not all hosts support all tests, which may lead to some tests to fail\n" +
			"                     for instance, ping packets may be dropped or the host does not support IPv6\n" +
			"                     specifying this will NOT limit assertions to ONLY this host\n\n" +
			"  --parallel       runs each test in parallel\n" +
			"  --no-col         do not colourize output\n\n" +
			"  --help           this help text\n" +
			"  --version        version\n",
	)
}

func version() {
	println("network tester v" + versionStr)
}
