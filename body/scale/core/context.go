// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package core

import (
	"bytes"
	"os"
	"os/user"

	"github.com/dawikur/dragon/body"
	"github.com/dawikur/dragon/config"
)

func username() string {
	user, _ := user.Current()
	return user.Username
}

func hostname() string {
	host, _ := os.Hostname()
	return host
}

func Context() body.Scale {
	return body.Scale{
		IsVisible: config.Core.Context.Visible,
		Color:     config.Core.Context.Color,
		RenderImpl: func(buffer *bytes.Buffer) {
			buffer.WriteString(username())
			buffer.WriteRune('@')
			buffer.WriteString(hostname())
		}}
}
