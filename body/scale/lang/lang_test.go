// Copyright 2018, Dawid Kurek, <dawikur@gmail.com>

package lang_test

import (
	"testing"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/body/scale/lang"
	"github.com/dawikur/dragon/test"
)

func TestLang_ScaleFunc(t *testing.T) {
	funcEmpty := func() string { return "" }
	funcNonEmpty := func() string { return "abcd" }

	for _, c := range []struct {
		description string
		mark        rune
		content     func() string
		expected    string
	}{
		{"Generated function will write only space if both mark and funcion are empty",
			' ',
			funcEmpty,
			" "},
		{"Generated function will write mark and space if they are not empty",
			'_',
			funcEmpty,
			"_ "},
		{"Generated function will write mark, space and content from function",
			'_',
			funcNonEmpty,
			"_ abcd"}} {
		expected := lang.Scale(false, body.Color{}, body.Mark{Content: c.mark, FG: body.None}, c.content)
		test.CheckRenderImpl(t, c.description, c.expected, expected.RenderImpl)
	}
}
