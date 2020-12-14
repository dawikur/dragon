// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package utils

import (
	"fmt"
)

type line struct {
	New  func()
	Up   func()
	Down func()
}

type term struct {
	Code  string
	Color string
	Reset func()
	Line  line
}

var (
	print = func(str string) func() {
		return func() { fmt.Print(str) }
	}

	// Term contains those fancy special strings which are used to control
	// terminal.
	Term = term{
		"%%{\u001b[%dm%%}",
		"%%{\u001b[%d;5;%dm%%}",
		print("%{\u001b[0m%}"),
		line{
			print("\n"),
			print("%{\u001b[1A%}"),
			print("%{\u001b[1B%}")}}
)
