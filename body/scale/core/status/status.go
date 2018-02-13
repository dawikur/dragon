// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package status

import (
	"strconv"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func getMark(retVal int) body.Mark {
	if retVal == 0 {
		return config.Core.Status.OK
	}
	if retVal >= 128 {
		return config.Core.Status.Signal
	}
	return config.Core.Status.Error
}

func getContent(retVal int) string {
	if retVal >= 128 {
		return strconv.Itoa(retVal - 128)
	}
	return strconv.Itoa(retVal)
}

func Scale(retVal int) body.Scale {
	return body.ScaleStr(
		retVal != 0,
		config.Core.Status.Color,
		getMark(retVal),
		getContent(retVal))
}
