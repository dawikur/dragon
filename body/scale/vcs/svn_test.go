// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package vcs_test

import (
	"bytes"
	"testing"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/body/scale/vcs"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/test"
	"github.com/dawikur/dragon/utils"
)

func TestSvn(t *testing.T) {
	cmd := utils.Cmd

	for _, c := range []struct {
		description string
		utilsCmds   []string
		isVisible   bool
		content     string
	}{
		{"If not svn repo don't render",
			[]string{""},
			false,
			""},
		{"Trunk clean branch",
			[]string{"^/trunk", ""},
			true,
			" trunk"},
		{"Clean branch different than trunk",
			[]string{"^/branch/other", ""},
			true,
			" other"},
		{"Trunk with untracked file",
			[]string{"^/trunk", "? file"},
			true,
			" trunk %{\u001b[38;5;4m%}"},
		{"Trunk with new file",
			[]string{"^/trunk", "A file"},
			true,
			" trunk %{\u001b[38;5;2m%}"},
		{"Trunk with deleted file",
			[]string{"^/trunk", "D file"},
			true,
			" trunk %{\u001b[38;5;1m%}"},
		{"Trunk with missing file",
			[]string{"^/trunk", "! file"},
			true,
			" trunk %{\u001b[38;5;1m%}"},
		{"Trunk with modified file",
			[]string{"^/trunk", "M file"},
			true,
			" trunk %{\u001b[38;5;17m%}"},
		{"Trunk with moved file",
			[]string{"^/trunk", "R file"},
			true,
			" trunk %{\u001b[38;5;6m%}"},
		{"Trunk with copied file",
			[]string{"^/trunk", "- file"},
			true,
			" trunk %{\u001b[38;5;3m%}"},
		{"Trunk with unmerged file",
			[]string{"^/trunk", "C file"},
			true,
			" trunk %{\u001b[38;5;5m%}═"},
		{"If branch doesn't have slash use whole name",
			[]string{"branch", "C file"},
			true,
			" branch %{\u001b[38;5;5m%}═"}} {

		utilsCmdIdx := -1
		utils.Cmd = func(string, ...string) string {
			utilsCmdIdx = utilsCmdIdx + 1
			return c.utilsCmds[utilsCmdIdx]
		}

		expected := body.Scale{
			IsVisible:  c.isVisible,
			Color:      config.Seg[2],
			RenderImpl: func(buffer *bytes.Buffer) { buffer.WriteString(c.content) }}

		scale := vcs.Svn()
		test.CheckScale(t, c.description, expected, scale)
	}

	utils.Cmd = cmd
}
