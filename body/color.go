// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package body

import (
	"bytes"
	"fmt"

	"github.com/dawikur/dragon/utils"
)

const (
	None      = -1
	Black     = 0
	Red       = 1
	Green     = 2
	Yellow    = 3
	Blue      = 4
	Violet    = 5
	Cyan      = 6
	LightGray = 7
	DarkGray  = 8
	White     = 15
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
	if color != None {
		fmt.Fprintf(buffer, utils.Term.Color, code, color)
	} else {
		fmt.Fprintf(buffer, utils.Term.Code, code+1)
	}
}
