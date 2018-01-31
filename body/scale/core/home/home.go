// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package home

import (
	"bytes"
	"os"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func Scale() body.Scale {
	dir, _ := os.Getwd()

	mark, color := func() (string, body.Color) {
		for _, c := range config.Core.Dir.SkipPrefixes {
			if strings.HasPrefix(dir, c.From) {
				return c.To, c.Color
			}
		}
		return "", body.Color{}
	}()

	return body.Scale{
		IsVisible: mark != "",
		Color:     color,
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(mark)
		}}
}
