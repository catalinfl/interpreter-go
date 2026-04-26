package parser

import (
	"fmt"
	"strings"
)

const traceEnabled = true
const traceIndentPlaceholder = "\t"

var traceLevel int

func identLevel() string {
	if traceLevel <= 0 {
		return ""
	}

	return strings.Repeat(traceIndentPlaceholder, traceLevel-1)
}

func tracePrint(msg string) {
	fmt.Printf("%s%s\n", identLevel(), msg)
}

func incIdent() {
	traceLevel++
}

func decIdent() {
	if traceLevel > 0 {
		traceLevel--
	}
}

func trace(msg string) string {
	if !traceEnabled {
		return msg
	}

	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	if !traceEnabled {
		return
	}

	tracePrint("END " + msg)
	decIdent()
}
