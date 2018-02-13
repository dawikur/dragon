// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package virtualenv

import (
	"os"
	"strings"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/body/scale/lang"
	"github.com/dawikur/dragon/config"
)

func Scale() body.Scale {
	env := os.Getenv("VIRTUAL_ENV")

	return lang.Scale(
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
