// Copyright 2017, Dawid Kurek, <dawikur@gmail.com>

// Package config contains all configuration for dragon
package config

import (
	"os"

	"github.com/dawikur/dragon/body"
)

type coreContext struct {
	Color   body.Color
	Visible bool
}

type coreDirSkipPrefix struct {
	Color body.Color
	From  string
	To    string
}

type coreDir struct {
	Color         body.Color
	Deepth        int
	MoreIndicator string
	JoinSeparator string
	SkipPrefixes  []coreDirSkipPrefix
}

type corePrompt struct {
	Color int
	Mark  string
}

type coreStatus struct {
	Color  body.Color
	OK     body.Mark
	Error  body.Mark
	Signal body.Mark
}

type coreSuspended struct {
	Color body.Color
}

type core struct {
	Context   coreContext
	Dir       coreDir
	Prompt    corePrompt
	Status    coreStatus
	Suspended coreSuspended
}

type lang struct {
	Color      body.Color
	Golang     body.Mark
	Nodejs     body.Mark
	Perl       body.Mark
	Php        body.Mark
	Python     body.Mark
	Ruby       body.Mark
	VirtualEnv body.Mark
}

type vcsBranch struct {
	Initial     body.Mark
	Tracked     body.Mark
	Detached    body.Mark
	Ahead       body.Mark
	Behind      body.Mark
	AheadBehind body.Mark
	Unknown     body.Mark
}

type vcsStatus struct {
	New      body.Mark
	Added    body.Mark
	Deleted  body.Mark
	Modified body.Mark
	Moved    body.Mark
	Copied   body.Mark
	Unmerged body.Mark
}

type vcs struct {
	Color  body.Color
	Branch vcsBranch
	Status vcsStatus
}

var (
	BG = []int{
		20,
		19,
		18,
		0}

	FG = 20

	Seg = []body.Color{
		{FG: body.Default, BG: BG[0]},
		{FG: FG, BG: BG[1]},
		{FG: FG, BG: BG[2]},
		{FG: FG, BG: BG[3]}}

	Core = core{
		Context: coreContext{
			Color:   Seg[0],
			Visible: os.Getenv("SSH_CONNECTION") != "",
		},
		Dir: coreDir{
			Color:         Seg[1],
			Deepth:        4,
			MoreIndicator: "…",
			JoinSeparator: "│",
			SkipPrefixes: []coreDirSkipPrefix{
				{Color: body.Color{FG: FG, BG: body.Bright}, From: os.Getenv("HOME"), To: ""},
				{Color: body.Color{FG: body.Default, BG: body.Blue}, From: "/tmp", To: ""},
				{Color: body.Color{FG: body.Default, BG: body.Orange}, From: "/", To: ""}},
		},
		Prompt: corePrompt{
			Color: body.Green,
			Mark:  "\n",
		},
		Status: coreStatus{
			Color:  Seg[3],
			OK:     body.Mark{},
			Error:  body.Mark{'', body.Red},
			Signal: body.Mark{'', body.Brown}},
		Suspended: coreSuspended{
			Color: Seg[0],
		},
	}

	Lang = lang{
		Color:      Seg[2],
		Golang:     body.Mark{'', body.Yellow},
		Nodejs:     body.Mark{'', body.Orange},
		Perl:       body.Mark{'', body.Blue},
		Php:        body.Mark{'', body.Magenta},
		Python:     body.Mark{'', body.Blue},
		Ruby:       body.Mark{'', body.Red},
		VirtualEnv: body.Mark{'', body.Blue},
	}

	VCS = vcs{
		Color: Seg[2],
		Branch: vcsBranch{
			Initial:     body.Mark{'', body.Magenta},
			Tracked:     body.Mark{'', body.None},
			Detached:    body.Mark{'', body.Yellow},
			Ahead:       body.Mark{'', body.Green},
			Behind:      body.Mark{'', body.Red},
			AheadBehind: body.Mark{'', body.Blue},
			Unknown:     body.Mark{'', body.Orange}},
		Status: vcsStatus{
			New:      body.Mark{'', body.Blue},
			Added:    body.Mark{'', body.Green},
			Deleted:  body.Mark{'', body.Red},
			Modified: body.Mark{'', body.Orange},
			Moved:    body.Mark{'', body.Cyan},
			Copied:   body.Mark{'', body.Yellow},
			Unmerged: body.Mark{'═', body.Magenta}},
	}
)
