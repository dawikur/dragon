// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package home

import (
	"bytes"
	"os"
	"strings"

	"../../../../body"
	"../../../../config"
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
		mark != "",
		color,
		func(buffer *bytes.Buffer) {
			buffer.WriteString(mark)
		}}
}
