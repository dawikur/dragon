// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package python

import (
	"bytes"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
	"github.com/dawikur/dragon/utils"
)

func getPythonVersion() string {
	ver := utils.Version("python", "--version")
	if len(ver) > 0 {
		return ver
	}

	buffer := bytes.Buffer{}
	separator := ""

	ver = utils.Version("python2", "--version")
	if len(ver) > 0 {
		separator = " "
		buffer.WriteString(ver)
	}

	ver = utils.Version("python3", "--version")
	if len(ver) > 0 {
		buffer.WriteString(separator)
		buffer.WriteString(ver)
	}

	return buffer.String()
}

func Scale() body.Scale {
	return body.ScaleFunc(
		utils.IsFile("[a-zA-Z]*.py"),
		config.Lang.Color,
		config.Lang.Python,
		getPythonVersion)
}
