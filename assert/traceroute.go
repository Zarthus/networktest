//go:build !windows

package assert

import (
	"errors"
	"github.com/aeden/traceroute"
	"github.com/zarthus/networktest/v2/assert/lib"
	"github.com/zarthus/networktest/v2/network"
	"strconv"
	"time"
)

func TracerouteV4(args lib.Args) lib.Result {
	host, _ := lib.SplitHostAndPort(lib.GetHost(args, network.W3OrgV4Address+":443"))
	res, err := runTraceroute(host, time.Duration(args.TimeoutSeconds)/time.Second)

	return lib.ResultByError("Traceroute using v4 IP", "traceroute: "+host, err, tracerouteGuidance(res))
}

func TracerouteV6(args lib.Args) lib.Result {
	host, _ := lib.SplitHostAndPort(lib.GetHost(args, network.W3OrgV6Address+":443"))
	res, err := runTraceroute(host, time.Duration(args.TimeoutSeconds)/time.Second)

	return lib.ResultByError("Traceroute using v6 IP", "traceroute: "+host, err, tracerouteGuidance(res))
}

func runTraceroute(host string, timeout time.Duration) (*traceroute.TracerouteResult, error) {
	opts := traceroute.TracerouteOptions{}
	opts.SetMaxHops(10)
	opts.SetRetries(2)
	opts.SetTimeoutMs(500) // TODO: don't ignore timeout val

	tracert, err := traceroute.Traceroute(host, &opts)

	if err == nil && len(tracert.Hops) == 0 {
		err = errors.New("traceroute with zero hops is likely indicative of an issue")
	}

	return &tracert, err
}

func tracerouteGuidance(res *traceroute.TracerouteResult) []string {
	if res == nil {
		return nil
	}

	var guidance []string
	for idx, hop := range res.Hops {
		guidance = append(guidance, "Hop "+strconv.Itoa(1+idx)+" ["+strconv.Itoa(int(hop.ElapsedTime.Milliseconds()))+" millis] "+hop.HostOrAddressString())
	}

	return guidance
}
