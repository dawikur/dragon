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
	skipPrefixes := []string{os.Getenv("HOME"), "/tmp"}

	for _, prefix := range skipPrefixes {
		if strings.HasPrefix(dir, prefix) {
			dir = dir[len(prefix):]
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

func getCurrentDir() string {
	dir, _ := os.Getwd()
	osSeparator := string(os.PathSeparator)

	dir = filterDirPrefixes(dir, osSeparator)
	dir = limitDirDeepth(dir, osSeparator)

	return dir
}

func Dir() body.Scale {
	currentDir := getCurrentDir()

	return body.Scale{
		IsVisible: currentDir != "",
		Color:     config.Core.Dir.Color,
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(currentDir)
		}}
}
