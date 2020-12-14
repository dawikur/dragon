// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package core

import (
	"bytes"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func Prompt() body.Scale {
	return body.Scale{
		IsVisible: true,
		Color:     body.Color{FG: config.Core.Prompt.Color, BG: body.None},
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(config.Core.Prompt.Mark)
		}}
}
