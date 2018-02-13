// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package git

import (
	"bytes"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/body/scale/vcs"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func parseStatus(info string) string {
	buffer := bytes.NewBufferString(" ")
	for _, check := range []struct {
		mark  body.Mark
		index string
		tree  string
	}{
		{config.VCS.Status.New, "\n?", "\n ?"},
		{config.VCS.Status.Added, "\nA", "\n A"},
		{config.VCS.Status.Deleted, "\nD", "\n D"},
		{config.VCS.Status.Modified, "\nM", "\n M"},
		{config.VCS.Status.Moved, "\nR", "\n R"},
		{config.VCS.Status.Copied, "\nC", "\n C"},
		{config.VCS.Status.Unmerged, "\nU", "\n U"}} {
		if strings.Contains(info, check.index) {
			check.mark.Render(buffer)
		} else if strings.Contains(info, check.tree) {
			check.mark.FG = config.VCS.Color.FG
			check.mark.Render(buffer)
		}
	}
	if buffer.Len() != 1 {
		return buffer.String()
	}
	return ""
}

func parseRepo(info string) body.Mark {
	value := 0
	for _, check := range []struct {
		id  int
		str string
	}{
		{1, "Initial commit on"},
		{2, "..."},
		{4, "(no branch)"},
		{8, "ahead"},
		{16, "behind"}} {
		if strings.Contains(info, check.str) {
			value += check.id
		}
	}

	switch value {
	case 1:
		return config.VCS.Branch.Initial
	case 2:
		return config.VCS.Branch.Tracked
	case 4:
		return config.VCS.Branch.Detached
	case 10:
		return config.VCS.Branch.Ahead
	case 18:
		return config.VCS.Branch.Behind
	case 26:
		return config.VCS.Branch.AheadBehind
	}
	return config.VCS.Branch.Unknown
}

func parseBranch(info string) string {
	info = info[3:]

	if strings.HasPrefix(info, "Initial commit on ") {
		parts := strings.Split(info, "\n")
		parts = strings.Split(parts[0], " ")
		return parts[3]
	}

	if strings.Contains(info, "(no branch)") {
		info = utils.Cmd("git", "rev-parse", "--short", "HEAD")
	}

	info = strings.Split(info, "\n")[0]
	info = strings.Split(info, "...")[0]
	return info
}

func Scale() body.Scale {
	return vcs.Scale(
		utils.Cmd("git", "status", "--porcelain", "--branch"),
		parseStatus, parseRepo, parseBranch)
}
