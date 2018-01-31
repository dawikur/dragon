// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package context

import (
	"bytes"
	"os"
	"os/user"

	"../../../../body"
	"../../../../config"
)

func username() string {
	user, _ := user.Current()
	return user.Username
}

func hostname() string {
	host, _ := os.Hostname()
	return host
}

func Scale() body.Scale {
	return body.Scale{
		config.Core.Context.Visible,
		config.Core.Context.Color,
		func(buffer *bytes.Buffer) {
			buffer.WriteString(username())
			buffer.WriteRune('@')
			buffer.WriteString(hostname())
		}}
}
