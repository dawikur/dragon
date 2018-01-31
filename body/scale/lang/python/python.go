// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package python

import (
	"bytes"

	"../../../../body"
	"../../../../config"
	"../../../../utils"
)

func Scale() body.Scale {
	return body.ScaleFunc(
		utils.IsFile("[a-zA-Z]*.py"),
		config.Lang.Color,
		config.Lang.Python,
		func() string {
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
		})
}
