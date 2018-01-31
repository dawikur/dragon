// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package nodejs

import (
	"../../../../body"
	"../../../../config"
	"../../../../utils"
)

func Scale() body.Scale {
	return body.ScaleFunc(
		utils.IsFile("[a-zA-Z]*.js"),
		config.Lang.Color,
		config.Lang.Nodejs,
		func() string {
			return utils.Version("nodejs", "--version")
		})
}
