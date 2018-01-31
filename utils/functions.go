// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	reVersion = regexp.MustCompile("[0-9]+.[0-9]+.[0-9]+")

	// Cmd is a wrapper for exec.Command which does not return error
	// but rather returns empty string if something went wrong
	Cmd = func(name string, args ...string) string {
		cmd := exec.Command(name, args...)
		ref, err := cmd.Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(ref))
	}

	// IsFile checks if in working directory, any file of a provided extension
	// actually do exists
	IsFile = func(ext string) bool {
		matches, err := filepath.Glob(ext)
		if err != nil {
			return false
		}
		return len(matches) > 0
	}

	// Version executes command (name arg) and tries to match its output
	// with a version regexp
	Version = func(name string, arg string) string {
		ver := Cmd(name, arg)
		return reVersion.FindString(ver)
	}

	// EnvOrValue returns environment variable or value of env is empty
	EnvOrValue = func(name string, value string) string {
		env := os.Getenv(name)
		if env != "" {
			return env
		}
		return value
	}
)
