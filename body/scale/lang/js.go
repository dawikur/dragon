// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package lang

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func getJSVersion() string {
	ver := utils.Version("nodejs", "--version")
	if len(ver) > 0 {
		return ver
	}

	ver = utils.Version("npm", "--version")
	if len(ver) > 0 {
		return ver
	}

	return ver
}

func Js() body.Scale {
	return Scale(
		utils.IsFile("[a-zA-Z]*.js"),
		config.Lang.Color,
		config.Lang.Nodejs,
		getJSVersion)
}
