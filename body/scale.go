// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package body

import (
	"bytes"
	"time"
)

type RenderFunc func(*bytes.Buffer)

type Scale struct {
	IsVisible  bool
	Color      Color
	RenderImpl RenderFunc
}

func ScaleStr(isVisible bool, color Color, mark Mark, content string) Scale {
	return ScaleFunc(isVisible, color, mark, func() string { return content })
}

func ScaleFunc(isVisible bool, color Color, mark Mark, content func() string) Scale {
	return Scale{
		isVisible,
		color,
		func(buffer *bytes.Buffer) {
			mark.Render(buffer)
			buffer.WriteRune(' ')
			buffer.WriteString(content())
		}}
}

func doRender(scale Scale, out chan<- string) {
	buffer := &bytes.Buffer{}

	scale.Color.Render(buffer)

	buffer.WriteRune(' ')
	scale.RenderImpl(buffer)
	buffer.WriteRune(' ')

	out <- buffer.String()
}

func (scale Scale) Render(out chan<- string) {
	if scale.IsVisible {
		result := make(chan string, 1)
		go doRender(scale, result)

		select {
		case result := <-result:
			out <- result
		case <-time.After(400 * time.Millisecond):
			break
		}
	}
	close(out)
}
