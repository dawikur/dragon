// Copyright 2018, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"bytes"

	"github.com/dawikur/dragon/body"
)

func Scale(isVisible bool, color body.Color, mark body.Mark, content func() string) body.Scale {
	return body.Scale{
		IsVisible: isVisible,
		Color:     color,
		RenderImpl: func(buffer *bytes.Buffer) {
			mark.Render(buffer)
			buffer.WriteRune(' ')
			buffer.WriteString(content())
		}}
}
