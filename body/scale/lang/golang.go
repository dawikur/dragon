// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func GoLang() body.Scale {
	return Scale(
		utils.IsFile("[a-zA-Z]*.go"),
		config.Lang.Color,
		config.Lang.Golang,
		func() string {
			return utils.Version("go", "version")
		})
}
