// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package prompt

import (
	"bytes"

	"../../../../body"
	"../../../../config"
)

func Scale() body.Scale {
	return body.Scale{
		true,
		body.Color{FG: config.Core.Prompt.Color},
		func(buffer *bytes.Buffer) {
			buffer.WriteString(config.Core.Prompt.Mark)
		}}
}
