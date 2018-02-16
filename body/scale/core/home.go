// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package core

import (
	"bytes"
	"os"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func getHomeMarkAndColor() (string, body.Color) {
	dir, _ := os.Getwd()

	for _, c := range config.Core.Dir.SkipPrefixes {
		if strings.HasPrefix(dir, c.From) {
			return c.To, c.Color
		}
	}
	return "", body.Color{}
}

func Home() body.Scale {
	mark, color := getHomeMarkAndColor()

	return body.Scale{
		IsVisible: mark != "",
		Color:     color,
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(mark)
		}}
}
