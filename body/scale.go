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
