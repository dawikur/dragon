// Copyright 2018, Dawid Kurek, <dawikur@gmail.com>

package vcs_test

import (
	"bytes"
	"testing"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/body/scale/vcs"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/test"
)

func TestCVS(t *testing.T) {

	for _, c := range []struct {
		description string
		info        string
		parseStatus func(string) string
		parseRepo   func(string) body.Mark
		parseBranch func(string) string
		isVisible   bool
		content     string
	}{
		{"If invisible don't render",
			"",
			func(string) string { return "" },
			func(string) body.Mark { return body.Mark{} },
			func(string) string { return "" },
			false,
			""}} {

		expected := body.Scale{
			IsVisible:  c.isVisible,
			Color:      config.Seg[2],
			RenderImpl: func(buffer *bytes.Buffer) { buffer.WriteString(c.content) }}

		scale := vcs.Scale(c.info, c.parseStatus, c.parseRepo, c.parseBranch)
		test.CheckScale(t, c.description, expected, scale)
	}
}
