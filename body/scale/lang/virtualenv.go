// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"os"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func VirtualEnv() body.Scale {
	env := os.Getenv("VIRTUAL_ENV")

	return Scale(
		env != "",
		config.Lang.Color,
		config.Lang.VirtualEnv,
		func() string {
			if parts := strings.Split(env, "/"); len(parts) > 0 {
				return parts[len(parts)-1]
			}
			return env
		})
}
