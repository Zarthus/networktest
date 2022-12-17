//go:build !windows

package main

import (
	"github.com/zarthus/networktest/v2/assert"
	"github.com/zarthus/networktest/v2/assert/lib"
)

func suiteExtras(args lib.Args) []item {
	return []item{
		{method: v4, Runnable: assert.TracerouteV4},
		{method: v6, Runnable: assert.TracerouteV6},
	}
}
