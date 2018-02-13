// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package perl

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/body/scale/lang"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func Scale() body.Scale {
	return lang.Scale(
		utils.IsFile("[a-zA-Z]*.pl"),
		config.Lang.Color,
		config.Lang.Perl,
		func() string {
			return utils.Version("perl", "--version")
		})
}
