// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package body_test

import (
	"testing"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/test"
)

func TestMark_Render(t *testing.T) {
	for _, c := range []struct {
		description string
		mark        body.Mark
		expected    string
	}{
		{"Invisible color will not be rendered",
			body.Mark{Content: '_', FG: body.None},
			"_"},
		{"Visible color will be printed",
			body.Mark{Content: 'a', FG: 1},
			"%{\u001b[38;5;1m%}a"}} {
		test.CheckRender(t, c.description, c.expected, c.mark)
	}
}
