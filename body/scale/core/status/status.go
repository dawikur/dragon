// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package status

import (
	"strconv"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func Scale(retval int) body.Scale {
	return body.ScaleStr(
		retval != 0,
		config.Core.Status.Color,
		func() body.Mark {
			if retval == 0 {
				return config.Core.Status.OK
			}
			if retval >= 128 {
				return config.Core.Status.Signal
			}
			return config.Core.Status.Error
		}(),
		func() string {
			if retval >= 128 {
				return strconv.Itoa(retval - 128)
			}
			return strconv.Itoa(retval)
		}())
}
