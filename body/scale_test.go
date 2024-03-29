// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package body_test

import (
	"bytes"
	"testing"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/test"
)

func TestScale_Render(t *testing.T) {
	renderEmpty := func(*bytes.Buffer) {}
	renderNonEmpty := func(buffer *bytes.Buffer) {
		buffer.WriteString("abcd")
	}

	for _, c := range []struct {
		description string
		scale       body.Scale
		expected    string
	}{
		{"Not visible scale should render nothing.",
			body.Scale{IsVisible: false},
			""},
		{"Visible scale with not visible color does not render color",
			body.Scale{IsVisible: true, Color: body.Color{FG: body.None, BG: body.None}, RenderImpl: renderEmpty},
			"  "},
		{"Visible scale with visible color of FG renders only FG color",
			body.Scale{IsVisible: true, Color: body.Color{FG: 1, BG: body.None}, RenderImpl: renderEmpty},
			"%{\u001b[38;5;1m%}  "},
		{"Visible scale with visible color of FG and BG renders only both colors",
			body.Scale{IsVisible: true, Color: body.Color{FG: 2, BG: 3}, RenderImpl: renderEmpty},
			"%{\u001b[38;5;2m%}%{\u001b[48;5;3m%}  "},
		{"Visible scale with visible color of FG and content renders FG and content",
			body.Scale{IsVisible: true, Color: body.Color{FG: 2, BG: body.None}, RenderImpl: renderNonEmpty},
			"%{\u001b[38;5;2m%} abcd "},
		{"Visible scale with visible color of FG and BG and content renders everything",
			body.Scale{IsVisible: true, Color: body.Color{FG: 2, BG: 11}, RenderImpl: renderNonEmpty},
			"%{\u001b[38;5;2m%}%{\u001b[48;5;11m%} abcd "}} {
		test.CheckRenderChan(t, c.description, c.expected, c.scale)
	}
}
