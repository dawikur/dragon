// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package perl

import (
	"../../../../body"
	"../../../../config"
	"../../../../utils"
)

func Scale() body.Scale {
	return body.ScaleFunc(
		utils.IsFile("[a-zA-Z]*.pl"),
		config.Lang.Color,
		config.Lang.Perl,
		func() string {
			return utils.Version("perl", "--version")
		})
}
