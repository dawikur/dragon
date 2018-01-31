// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package php

import (
	"../../../../body"
	"../../../../config"
	"../../../../utils"
)

func Scale() body.Scale {
	return body.ScaleFunc(
		utils.IsFile("[a-zA-Z]*.php"),
		config.Lang.Color,
		config.Lang.Php,
		func() string {
			return utils.Version("php", "--version")
		})
}
