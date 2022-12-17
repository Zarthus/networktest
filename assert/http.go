package assert

import (
	"github.com/zarthus/networktest/v2/assert/lib"
	"github.com/zarthus/networktest/v2/network"
	"net/http"
)

func QueryHttpsDns(args lib.Args) lib.Result {
	host, _ := lib.SplitHostAndPort(lib.GetHost(args, "w3.org"))
	status, err := queryHttp("https://", host)

	return lib.ResultByError("HTTPS using DNS", lib.ConnectionAssertion(host, 443), err, []string{"HTTP Status Code: " + status})
}

func QueryHttpV4(args lib.Args) lib.Result {
	host, _ := lib.SplitHostAndPort(lib.GetHost(args, network.W3OrgV4Address))
	status, err := queryHttp("http://", host)

	return lib.ResultByError("HTTP v4 IP", lib.ConnectionAssertion(host, 80), err, []string{"HTTP Status Code: " + status})
}

func QueryHttpV6(args lib.Args) lib.Result {
	host, _ := lib.SplitHostAndPort(lib.GetHost(args, network.W3OrgV6Address))
	status, err := queryHttp("http://", host)

	return lib.ResultByError("HTTP v6 IP", lib.ConnectionAssertion(host, 80), err, []string{"HTTP Status Code: " + status})
}

func queryHttp(scheme string, host string) (string, error) {
	response, err := http.Get(scheme + host)
	if err != nil {
		return "N/A", err
	}

	defer response.Body.Close()
	return response.Status, nil
}
