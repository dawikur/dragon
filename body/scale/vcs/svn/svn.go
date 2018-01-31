// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package svn

import (
	"bytes"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func Scale() body.Scale {
	mark := body.Mark{}
	status := ""
	branch := ""

	parseStatus := func(info string) string {
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

	parseRepo := func(info string) body.Mark {
		return config.VCS.Branch.Tracked
	}

	parseBranch := func(info string) string {
		parts := strings.Split(info, "/")
		if len(parts) < 2 {
			return info
		}

		if parts[1] == "trunk" {
			return "trunk"
		}

		return parts[2]
	}

	checkRepo := func() bool {
		info := utils.Cmd("svn", "info", "--show-item", "relative-url", "--no-newline")
		if info == "" {
			return false
		}

		status = parseStatus(info)
		mark = parseRepo(info)
		branch = parseBranch(info)

		return true
	}

	return body.Scale{
		IsVisible: checkRepo(),
		Color:     config.VCS.Color,
		RenderImpl: func(buffer *bytes.Buffer) {
			mark.Render(buffer)
			buffer.WriteRune(' ')
			buffer.WriteString(branch)
			buffer.WriteString(status)
		}}
}
