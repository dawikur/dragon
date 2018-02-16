// Copyright 2018, Dawid Kurek, <dawikur@gmail.com>

package main

import (
	"github.com/dawikur/dragon/app"

	"github.com/dawikur/dragon/body/scale/core/context"
	"github.com/dawikur/dragon/body/scale/core/dir"
	"github.com/dawikur/dragon/body/scale/core/home"
	"github.com/dawikur/dragon/body/scale/core/prompt"
	"github.com/dawikur/dragon/body/scale/core/status"
	"github.com/dawikur/dragon/body/scale/core/suspended"

	"github.com/dawikur/dragon/body/scale/lang"

	"github.com/dawikur/dragon/body/scale/vcs/git"
	"github.com/dawikur/dragon/body/scale/vcs/svn"

	"github.com/dawikur/dragon/utils"
)

func main() {
	utils.Term.Line.New()

	utils.Term.Line.Up()
	defer utils.Term.Line.Down()

	app.Run(
		// core
		context.Scale(),
		dir.Scale(),
		home.Scale(),
		prompt.Scale(),
		status.Scale(255),
		suspended.Scale("3"),

		// lang
		lang.Elm(),
		lang.GoLang(),
		lang.Js(),
		lang.Perl(),
		lang.Php(),
		lang.Python(),
		lang.Ruby(),
		lang.VirtualEnv(),

		// vcs
		git.Scale(),
		svn.Scale())
}
