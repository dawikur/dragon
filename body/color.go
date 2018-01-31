// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package body

import (
	"bytes"
	"fmt"

	"../utils"
)

const (
	None    = -1
	Default = 0
	Red     = 1
	Green   = 2
	Yellow  = 3
	Blue    = 4
	Magenta = 5
	Cyan    = 6
	White   = 7
	Bright  = 8
	Brown   = 16
	Orange  = 17
)

// Color gathers both FG and BG color codes
type Color struct {
	FG int
	BG int
}

// Render handles the magic-terminal sequence for changing
// color for FG and BG (if set)
func (color Color) Render(buffer *bytes.Buffer) {
	render(buffer, 38, color.FG)
	render(buffer, 48, color.BG)
}

func render(buffer *bytes.Buffer, code int, color int) {
	if color != -1 {
		fmt.Fprintf(buffer, utils.Term.Color, code, color)
	}
}
