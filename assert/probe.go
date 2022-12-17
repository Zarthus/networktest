package assert

import (
	"github.com/zarthus/networktest/v2/assert/lib"
	"github.com/zarthus/networktest/v2/network"
	"net"
	"strconv"
	"time"
)

// A series of very light checks, just checking if a connection gets accepted on a specific port

func ProbeDns(args lib.Args) lib.Result {
	host, port := lib.SplitHostAndPort(lib.GetHost(args, "w3.org:443"))
	err := probe(host, port, args.TimeoutSeconds)

	return lib.ResultByError("Probe using DNS", probeAssertion(host, port), err, nil)
}

func ProbeV4(args lib.Args) lib.Result {
	host, port := lib.SplitHostAndPort(lib.GetHost(args, network.W3OrgV4Address+":443"))
	err := probe(host, port, args.TimeoutSeconds)

	return lib.ResultByError("Probe using v4 IP", probeAssertion(host, port), err, nil)
}

func ProbeV6(args lib.Args) lib.Result {
	host, port := lib.SplitHostAndPort(lib.GetHost(args, network.W3OrgV6Address+":443"))
	err := probe(host, port, args.TimeoutSeconds)

	return lib.ResultByError("Probe using v6 IP", probeAssertion(host, port), err, nil)
}

func probe(host string, port int, argTimeout int) error {
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), time.Duration(argTimeout)*time.Second)

	if err != nil {
		return err
	}

	err2 := conn.Close()
	if err2 != nil {
		return err2
	}

	return nil
}

func probeAssertion(host string, port int) string {
	return "Connection to " + host + ":" + strconv.Itoa(port)
}
