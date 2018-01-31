// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

// Package app provides common entry point.
//
// It is enought to pass an array of body.Scale as argument
// and both executing&rendering will be handled internally.
package app

import (
	"fmt"
	"runtime"

	"github.com/dawikur/dragon/utils"
)

func Run(scales ...interface {
	Render(chan<- string)
}) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	chans := make([]chan string, len(scales))
	for idx := range chans {
		chans[idx] = make(chan string, 1)
	}

	for idx, scale := range scales {
		go scale.Render(chans[idx])
	}

	for _, c := range chans {
		fmt.Print(<-c)
	}

	utils.Term.Reset()
}
