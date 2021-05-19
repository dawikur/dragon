// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package core

import (
	"bytes"
	"os"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func filterDirPrefixes(dir string, osSeparator string) string {
	for _, prefix := range config.Core.Dir.SkipPrefixes {
		if strings.HasPrefix(dir, prefix.From) {
			dir = dir[len(prefix.From):]
			break
		}
	}

	if strings.HasPrefix(dir, osSeparator) {
		dir = dir[len(osSeparator):]
	}

	return dir
}

func limitDirDeepth(dir string, osSeparator string) string {
	dirs := strings.Split(dir, osSeparator)

	if len(dirs) > config.Core.Dir.Deepth {
		dirs = dirs[len(dirs)-config.Core.Dir.Deepth:]
		dirs[0] = config.Core.Dir.MoreIndicator
	}

	return strings.Join(dirs, config.Core.Dir.JoinSeparator)
}

func getCurrentDir() (string, body.Mark) {
	status := config.Core.Dir.Status.OK

	dir, err := os.Getwd()
	if err != nil {
		status = config.Core.Dir.Status.Error
		dir = os.Getenv("PWD")
	}
	osSeparator := string(os.PathSeparator)

	dir = filterDirPrefixes(dir, osSeparator)
	dir = limitDirDeepth(dir, osSeparator)

	return dir, status
}

func Dir() body.Scale {
	currentDir, status := getCurrentDir()

	return body.Scale{
		IsVisible: currentDir != "",
		Color:     config.Core.Dir.Color,
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(currentDir)
			status.Render(buffer)
		}}
}
