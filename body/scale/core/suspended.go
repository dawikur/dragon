// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package core

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func Suspended(jobs string) body.Scale {
	return body.ScaleStr(
		jobs != "0",
		config.Core.Suspended.Color,
		config.Core.Suspended.Mark,
		jobs)
}
