// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package vcs

import (
	"bytes"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func parseSvnStatus(info string) string {
	info = "\n" + utils.Cmd("svn", "status", "--ignore-externals")

	buffer := bytes.NewBufferString(" ")
	for _, check := range []struct {
		mark  body.Mark
		index string
	}{
		{config.VCS.Status.New, "\n?"},
		{config.VCS.Status.Added, "\nA"},
		{config.VCS.Status.Deleted, "\nD"},
		{config.VCS.Status.Deleted, "\n!"},
		{config.VCS.Status.Modified, "\nM"},
		{config.VCS.Status.Moved, "\nR"},
		{config.VCS.Status.Copied, "\n-"},
		{config.VCS.Status.Unmerged, "\nC"}} {
		if strings.Contains(info, check.index) {
			check.mark.Render(buffer)
		}
	}
	if buffer.Len() != 1 {
		return buffer.String()
	}
	return ""
}

func parseSvnRepo(info string) body.Mark {
	return config.VCS.Branch.Tracked
}

func parseSvnBranch(info string) string {
	parts := strings.Split(info, "/")
	if len(parts) < 2 {
		return info
	}

	if parts[1] == "trunk" {
		return "trunk"
	}

	return parts[2]
}

func Svn() body.Scale {
	return Scale(
		utils.Cmd("svn", "info", "--show-item", "relative-url", "--no-newline"),
		parseSvnStatus, parseSvnRepo, parseSvnBranch)
}
