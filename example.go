// Copyright 2018, Dawid Kurek, <dawikur@gmail.com>

package main

import (
	"github.com/dawikur/dragon/app"

	"github.com/dawikur/dragon/body/scale/core"
	"github.com/dawikur/dragon/body/scale/lang"
	"github.com/dawikur/dragon/body/scale/vcs"

	"github.com/dawikur/dragon/utils"
)

func main() {
	utils.Term.Line.New()

	utils.Term.Line.Up()
	defer utils.Term.Line.Down()

	app.Run(
		core.Context(),
		core.Dir(),
		core.Home(),
		core.Prompt(),
		core.Status(255),
		core.Suspended("3"),

		lang.Elm(),
		lang.GoLang(),
		lang.Js(),
		lang.Perl(),
		lang.Php(),
		lang.Python(),
		lang.Ruby(),
		lang.VirtualEnv(),

		vcs.Git(),
		vcs.Svn())
}
