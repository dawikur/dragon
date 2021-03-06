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

func TestGit(t *testing.T) {
	cmd := utils.Cmd

	for _, c := range []struct {
		description string
		utilsCmds   []string
		isVisible   bool
		content     string
	}{
		{"If not git repo don't render",
			[]string{""},
			false,
			""},
		{"Master clean branch",
			[]string{"## master...origin/master"},
			true,
			" master"},
		{"Master ahead branch",
			[]string{"## master...origin/master [ahead 2]"},
			true,
			"%{\u001b[38;5;2m%} master"},
		{"Master ahead and behind branch",
			[]string{"## master...origin/master [ahead 2, behind 3]"},
			true,
			"%{\u001b[38;5;5m%} master"},
		{"Master behind branch",
			[]string{"## master...origin/master [behind 2]"},
			true,
			"%{\u001b[38;5;4m%} master"},
		{"HEAD detached branch",
			[]string{"## HEAD (no branch)", "123456"},
			true,
			"%{\u001b[38;5;3m%} 123456"},
		{"Master with untracked file",
			[]string{"## master...origin/master\n?? file"},
			true,
			" master %{\u001b[38;5;7m%}"},
		{"Master with new file",
			[]string{"## master...origin/master\nA  file"},
			true,
			" master %{\u001b[38;5;2m%}"},
		{"Master with tracked new file",
			[]string{"## master...origin/master\n A file"},
			true,
			" master %{\u001b[38;5;7m%}"},
		{"Master with deleted file",
			[]string{"## master...origin/master\nD  file"},
			true,
			" master %{\u001b[38;5;1m%}"},
		{"Master with tracked deleted file",
			[]string{"## master...origin/master\n D file"},
			true,
			" master %{\u001b[38;5;7m%}"},
		{"Master with modified file",
			[]string{"## master...origin/master\nM  file"},
			true,
			" master %{\u001b[38;5;3m%}"},
		{"Master with tracked modified file",
			[]string{"## master...origin/master\n M file"},
			true,
			" master %{\u001b[38;5;7m%}"},
		{"Master with moved file",
			[]string{"## master...origin/master\nR  file"},
			true,
			" master %{\u001b[38;5;4m%}"},
		{"Master with tracked moved file",
			[]string{"## master...origin/master\n R file"},
			true,
			" master %{\u001b[38;5;7m%}"},
		{"Master with copied file",
			[]string{"## master...origin/master\nC  file"},
			true,
			" master %{\u001b[38;5;6m%}"},
		{"Master with tracked copied file",
			[]string{"## master...origin/master\n C file"},
			true,
			" master %{\u001b[38;5;7m%}"},
		{"Master with unmerged file",
			[]string{"## master...origin/master\nU  file"},
			true,
			" master %{\u001b[38;5;5m%}═"},
		{"Master with tracked unmerged file",
			[]string{"## master...origin/master\n U file"},
			true,
			" master %{\u001b[38;5;7m%}═"},
		{"Initial commit on something",
			[]string{"## Initial commit on something\n"},
			true,
			"%{\u001b[38;5;7m%} something"},
		{"Branch without remote",
			[]string{"## Branch\n"},
			true,
			"%{\u001b[38;5;6m%} Branch"},
		{"Branch with dot in name",
			[]string{"## Branch1.5...origin/Branch1.5\n"},
			true,
			" Branch1.5"}} {

		utilsCmdIdx := -1
		utils.Cmd = func(string, ...string) string {
			utilsCmdIdx = utilsCmdIdx + 1
			return c.utilsCmds[utilsCmdIdx]
		}

		expected := body.Scale{
			IsVisible:  c.isVisible,
			Color:      config.Seg[3],
			RenderImpl: func(buffer *bytes.Buffer) { buffer.WriteString(c.content) }}

		scale := vcs.Git()
		test.CheckScale(t, c.description, expected, scale)
	}

	utils.Cmd = cmd
}
