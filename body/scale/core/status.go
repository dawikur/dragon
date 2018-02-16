// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package core

import (
	"strconv"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func getStatusMark(retVal int) body.Mark {
	if retVal == 0 {
		return config.Core.Status.OK
	}
	if retVal >= 128 {
		return config.Core.Status.Signal
	}
	return config.Core.Status.Error
}

func getStatusContent(retVal int) string {
	if retVal >= 128 {
		return strconv.Itoa(retVal - 128)
	}
	return strconv.Itoa(retVal)
}

func Status(retVal int) body.Scale {
	return body.ScaleStr(
		retVal != 0,
		config.Core.Status.Color,
		getStatusMark(retVal),
		getStatusContent(retVal))
}
