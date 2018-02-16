// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func Elm() body.Scale {
	return Scale(
		utils.IsFile("[a-zA-Z]*.elm"),
		config.Lang.Color,
		config.Lang.Elm,
		func() string {
			return utils.Version("elm", "--version")
		})
}
