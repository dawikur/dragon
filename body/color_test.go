// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package body_test

import (
	"testing"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/test"
)

func TestColor_Render(t *testing.T) {
	for _, c := range []struct {
		description string
		color       body.Color
		expected    string
	}{
		{"Invisible color will not be rendered",
			body.Color{FG: body.None, BG: body.None},
			""},
		{"Color with only FG will render only FG",
			body.Color{FG: 8, BG: body.None},
			"%{\u001b[38;5;8m%}"},
		{"Color with only BG will render only BG",
			body.Color{FG: body.None, BG: 7},
			"%{\u001b[48;5;7m%}"},
		{"Color with FG and BG will render only both",
			body.Color{FG: 8, BG: 7},
			"%{\u001b[38;5;8m%}%{\u001b[48;5;7m%}"}} {
		test.CheckRender(t, c.description, c.expected, c.color)
	}
}
