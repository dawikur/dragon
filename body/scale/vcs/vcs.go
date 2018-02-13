// Copyright 2018, Dawid Kurek, <dawikur@gmail.com>

package vcs

import (
	"bytes"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func checkRepo(
	info string,
	parseStatus func(string) string,
	parseRepo func(string) body.Mark,
	parseBranch func(string) string) (bool, string, body.Mark, string) {
	if info == "" {
		return false, "", body.Mark{}, ""
	}

	return true,
		parseStatus(info),
		parseRepo(info),
		parseBranch(info)
}

func Scale(info string,
	parseStatus func(string) string,
	parseRepo func(string) body.Mark,
	parseBranch func(string) string) body.Scale {

	inRepo, status, mark, branch := checkRepo(
		info,
		parseStatus,
		parseRepo,
		parseBranch)

	return body.Scale{
		IsVisible: inRepo,
		Color:     config.VCS.Color,
		RenderImpl: func(buffer *bytes.Buffer) {
			mark.Render(buffer)
			buffer.WriteRune(' ')
			buffer.WriteString(branch)
			buffer.WriteString(status)
		}}
}
