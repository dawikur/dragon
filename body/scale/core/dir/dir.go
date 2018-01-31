// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package dir

import (
	"bytes"
	"os"
	"strings"

	"../../../../body"
	"../../../../config"
)

var (
	skipPrefixes = []string{os.Getenv("HOME"), "/tmp"}
)

func Scale() body.Scale {
	currentDir := func() string {
		dir, _ := os.Getwd()
		osSeparator := string(os.PathSeparator)

		for _, prefix := range skipPrefixes {
			if strings.HasPrefix(dir, prefix) {
				dir = dir[len(prefix):]
				break
			}
		}

		if strings.HasPrefix(dir, osSeparator) {
			dir = dir[1:]
		}

		dirs := strings.Split(dir, osSeparator)
		if len(dirs) > config.Core.Dir.Deepth {
			dirs = dirs[len(dirs)-config.Core.Dir.Deepth:]
			dirs[0] = config.Core.Dir.MoreIndicator
		}
		dir = strings.Join(dirs, config.Core.Dir.JoinSeparator)

		return dir
	}()

	return body.Scale{
		currentDir != "",
		config.Core.Dir.Color,
		func(buffer *bytes.Buffer) {
			buffer.WriteString(currentDir)
		}}
}
