package main

import (
	"github.com/zarthus/networktest/v2/assert/lib"
)

const (
	red   string = "31"
	green string = "32"
	amber string = "33"
)

func Write(result lib.Result, colourize bool) {
	switch result.Status {
	case lib.Ok:
		ok(result, colourize)
	case lib.Warn:
		warning(result, colourize)
	case lib.NotExecuted:
		warning(result, colourize)
	case lib.Error:
		danger(result, colourize)
	}
}

func ok(result lib.Result, colourize bool) {
	println(formatText("   OK ", result, green, colourize))
}

func warning(result lib.Result, colourize bool) {
	println(formatText(" WARN ", result, amber, colourize))
}

func danger(result lib.Result, colourize bool) {
	println(formatText("ERROR ", result, red, colourize))
}

func formatText(prefix string, result lib.Result, seq string, colourize bool) string {
	str := prefix + wrapColour(result.Check, seq, colourize) + " | " + result.Assertion + "\n"

	if len(result.Guidance) == 0 {
		return str
	}

	sym := "┠"
	for idx, guidance := range result.Guidance {
		if idx+1 == len(result.Guidance) {
			sym = "┗"
		}

		str = str + "      " + sym + " " + guidance + "\n"
	}

	return str
}

func wrapColour(s string, seq string, colourize bool) string {
	if !colourize {
		return s
	}

	return "\033[" + seq + "m" + s + "\033[0m"
}
