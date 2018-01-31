// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

// Package test provides testing function functions.
package test

import (
	"bytes"
	"strings"
	"testing"
)

type renderer interface {
	Render(chan<- string)
}

func check(t *testing.T, description string, expected string, result string) {
	if result != expected {
		t.Errorf(description+
			"\nexpected: \"%s\""+
			"\nresult:   \"%s\"",
			strings.Replace(expected, "\u001b", "\\u001b", -1),
			strings.Replace(result, "\u001b", "\\u001b", -1))
	}
}

func render(scale renderer) string {
	channel := make(chan string, 1)
	scale.Render(channel)
	result := <-channel
	return result
}

func CheckScale(t *testing.T, description string, expected renderer, result renderer) {
	check(t, description, render(expected), render(result))
}

func CheckRender(t *testing.T, description string, expected string, scale interface {
	Render(*bytes.Buffer)
}) {
	buffer := bytes.Buffer{}
	scale.Render(&buffer)
	result := buffer.String()
	check(t, description, expected, result)
}

func CheckRenderImpl(t *testing.T, description string, expected string, renderImpl func(*bytes.Buffer)) {
	buffer := bytes.Buffer{}
	renderImpl(&buffer)
	result := buffer.String()
	check(t, description, expected, result)
}

func CheckRenderChan(t *testing.T, description string, expected string, scale renderer) {
	result := render(scale)
	check(t, description, expected, result)
}
