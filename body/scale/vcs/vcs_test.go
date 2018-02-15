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
		mark        body.Mark
		isVisible   bool
		content     string
	}{
		{"If invisible don't render",
			"",
			body.Mark{},
			false,
			""},
		{"If visible render simple text",
			"Repo",
			body.Mark{Content: ' '},
			true,
			" Repo"}} {

		expected := body.Scale{
			IsVisible:  c.isVisible,
			Color:      config.Seg[2],
			RenderImpl: func(buffer *bytes.Buffer) { buffer.WriteString(c.content) }}

		scale := vcs.Scale(
			c.info,
			func(string) string { return "" },
			func(string) body.Mark { return c.mark },
			func(string) string { return c.info })
		test.CheckScale(t, c.description, expected, scale)
	}
}
