// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func Ruby() body.Scale {
	return Scale(
		utils.IsFile("[a-zA-Z]*.rb"),
		config.Lang.Color,
		config.Lang.Ruby,
		func() string {
			return utils.Version("ruby", "--version")
		})
}
