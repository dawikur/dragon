// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func Php() body.Scale {
	return Scale(
		utils.IsFile("[a-zA-Z]*.php"),
		config.Lang.Color,
		config.Lang.Php,
		func() string {
			return utils.Version("php", "--version")
		})
}
