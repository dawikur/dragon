// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package core

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func Suspended(jobs string) body.Scale {
	mark := config.Core.Suspended.Mark
	content := jobs

	return body.Scale{
		jobs != "0",
		config.Core.Suspended.Color,
		func(buffer *bytes.Buffer) {
			mark.Render(buffer)
			buffer.WriteRune(' ')
			buffer.WriteString(content)
		}}
}
