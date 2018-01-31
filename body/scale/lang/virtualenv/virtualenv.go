// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package virtualenv

import (
	"os"
	"strings"

	"../../../../body"
	"../../../../config"
)

func Scale() body.Scale {
	env := os.Getenv("VIRTUAL_ENV")

	return body.ScaleFunc(
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
