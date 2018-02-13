// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package dir

import (
	"bytes"
	"os"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func filterPrefixes(dir string, osSeparator string) string {
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

func limitDeepth(dir string, osSeparator string) string {
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

	dir = filterPrefixes(dir, osSeparator)
	dir = limitDeepth(dir, osSeparator)

	return dir
}

func Scale() body.Scale {
	currentDir := getCurrentDir()

	return body.Scale{
		IsVisible: currentDir != "",
		Color:     config.Core.Dir.Color,
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(currentDir)
		}}
}
