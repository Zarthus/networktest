# networktest

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/Zarthus/networktest.svg)](https://pkg.go.dev/github.com/Zarthus/networktest/v2)

A simple go binary that runs simple network connectivity tests

## Command Line Interface
```
network tester v0.1 - tests network connectivity

  --only=4|6       only test ipv4 or ipv6 connectivity [default both]
  --timeout=int    timeout in seconds [default 3]
  --target=string  runs some testcases against a specific hostname [default "auto"]
                     please be aware that not all hosts support all tests, which may lead to some tests to fail
                     for instance, ping packets may be dropped or the host does not support IPv6
                     specifying this will NOT limit assertions to ONLY this host

  --parallel       runs each test in parallel
  --no-col         do not colourize output

  --help           this help text
  --version        version
```


### Example output

```bash
$ go build && ./networktest

network tester v0.1
 WARN Assigned IP(s) | Check local IP configuration
      ┠ (OK) 123.123.123.123
      ┗ Missing IPv6 connectivity

   OK DNS Lookup (A / AAAA) | Resolve DNS of w3.org
      ┠ 128.30.52.100
      ┗ 2603:400a:ffff:804:801e:34:0:64

   OK DNS Lookup (MX) | Resolve DNS of w3.org
      ┠ mimas.w3.org.
      ┠ titan.w3.org.
      ┗ bart.w3.org.

   OK Probe using DNS | Connection to w3.org:443

   OK Probe using v4 IP | Connection to 128.30.52.100:443

ERROR Probe using v6 IP | Connection to [2603:400a:ffff:804:801e:34:0:64]:443
      ┗ dial tcp [2603:400a:ffff:804:801e:34:0:64]:443: connect: network is unreachable

   OK Traceroute using v4 IP | traceroute: 128.30.52.100
      ┠ Hop 1 [0 millis] local.
      ┠ Hop 2 [0 millis] modem.
      ┠ Hop 3 [2 millis] kpn.
      ┠ Hop 4 [5 millis] amsterdam.
      ┠ Hop 5 [5 millis] nl-ams04a-ri3-ae-9-0.aorta.net.
      ┠ Hop 6 [4 millis] ae254.border-a.sech-ams.netarch.akamai.com.
      ┗ Hop 7 [5 millis] po110.bs-b.sech-ams.netarch.akamai.com.

ERROR Traceroute using v6 IP | traceroute: [2603:400a:ffff:804:801e:34:0:64]
      ┗ traceroute with zero hops is likely indicative of an issue
```
