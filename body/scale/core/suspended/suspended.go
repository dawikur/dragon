// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package suspended

import (
	"../../../../body"
	"../../../../config"
)

func Scale(jobs string) body.Scale {
	return body.ScaleStr(
		jobs != "0",
		config.Core.Suspended.Color,
		body.Mark{'', config.Core.Suspended.Color.FG},
		jobs)
}
