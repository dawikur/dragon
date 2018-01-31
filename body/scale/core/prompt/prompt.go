// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package prompt

import (
	"bytes"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func Scale() body.Scale {
	return body.Scale{
		IsVisible: true,
		Color:     body.Color{FG: config.Core.Prompt.Color},
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(config.Core.Prompt.Mark)
		}}
}
