// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package suspended

import (
	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func Scale(jobs string) body.Scale {
	return body.ScaleStr(
		jobs != "0",
		config.Core.Suspended.Color,
		body.Mark{Content: '', FG: config.Core.Suspended.Color.FG},
		jobs)
}
