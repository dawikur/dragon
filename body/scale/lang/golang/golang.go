// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package golang

import (
	"../../../../body"
	"../../../../config"
	"../../../../utils"
)

func Scale() body.Scale {
	return body.ScaleFunc(
		utils.IsFile("[a-zA-Z]*.go"),
		config.Lang.Color,
		config.Lang.Golang,
		func() string {
			return utils.Version("go", "version")
		})
}
