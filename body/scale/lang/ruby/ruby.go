// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package ruby

import (
	"../../../../body"
	"../../../../config"
	"../../../../utils"
)

func Scale() body.Scale {
	return body.ScaleFunc(
		utils.IsFile("[a-zA-Z]*.rb"),
		config.Lang.Color,
		config.Lang.Ruby,
		func() string {
			return utils.Version("ruby", "--version")
		})
}
