// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package body

import (
	"bytes"
)

type Mark struct {
	Content rune
	FG      int
}

func (mark Mark) Render(buffer *bytes.Buffer) {
	if mark.Content != ' ' {
		color := Color{FG: mark.FG, BG: -1}
		color.Render(buffer)
		buffer.WriteRune(mark.Content)
	}
}
