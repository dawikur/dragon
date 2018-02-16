// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func Perl() body.Scale {
	return Scale(
		utils.IsFile("[a-zA-Z]*.pl"),
		config.Lang.Color,
		config.Lang.Perl,
		func() string {
			return utils.Version("perl", "--version")
		})
}
