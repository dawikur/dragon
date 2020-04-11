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
	Mark  body.Mark
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
	Elm        body.Mark
	Golang     body.Mark
	Nodejs     body.Mark
	Perl       body.Mark
	Php        body.Mark
	Python     body.Mark
	Ruby       body.Mark
	VirtualEnv body.Mark
}

type vcsBranch struct {
	Empty       body.Mark
	Initial     body.Mark
	Tracked     body.Mark
	Detached    body.Mark
	Ahead       body.Mark
	Behind      body.Mark
	AheadBehind body.Mark
	Fatal       body.Mark
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
	Split  string
	Join   string
	Branch vcsBranch
	Status vcsStatus
}

var (
	Default = body.Default

	BG = []int{
		body.White,
		body.LightGray,
		body.DarkGray,
		body.Black}

	FG = []int{
		Default,
		Default,
		body.LightGray,
		body.LightGray}

	Seg = []body.Color{
		{FG: FG[0], BG: BG[0]},
		{FG: FG[1], BG: BG[1]},
		{FG: FG[2], BG: BG[2]},
		{FG: FG[3], BG: BG[3]}}

	Core = core{
		Context: coreContext{
			Color:   Seg[0],
			Visible: os.Getenv("SSH_CONNECTION") != "",
		},
		Dir: coreDir{
			Color:         Seg[2],
			Deepth:        4,
			MoreIndicator: "…",
			JoinSeparator: "│",
			SkipPrefixes: []coreDirSkipPrefix{
				{Color: body.Color{FG: Default, BG: body.LightGray}, From: os.Getenv("HOME"), To: ""},
				{Color: body.Color{FG: Default, BG: body.Blue}, From: "/tmp", To: ""},
				{Color: body.Color{FG: Default, BG: body.Violet}, From: "/mnt", To: ""},
				{Color: body.Color{FG: Default, BG: body.Yellow}, From: "/media", To: ""},
				{Color: body.Color{FG: Default, BG: body.Red}, From: "/", To: ""}},
		},
		Prompt: corePrompt{
			Color: body.Green,
			Mark:  "\n",
		},
		Status: coreStatus{
			Color:  Seg[3],
			OK:     body.Mark{},
			Error:  body.Mark{Content: '', FG: body.Red},
			Signal: body.Mark{Content: '', FG: body.Yellow}},
		Suspended: coreSuspended{
			Color: Seg[1],
			Mark:  body.Mark{Content: '', FG: Default},
		},
	}

	Lang = lang{
		Color:      Seg[2],
		Elm:        body.Mark{Content: '', FG: body.Green},
		Golang:     body.Mark{Content: '', FG: body.Blue},
		Nodejs:     body.Mark{Content: '', FG: body.Yellow},
		Perl:       body.Mark{Content: '', FG: body.Blue},
		Php:        body.Mark{Content: '', FG: body.Violet},
		Python:     body.Mark{Content: '', FG: body.Blue},
		Ruby:       body.Mark{Content: '', FG: body.Red},
		VirtualEnv: body.Mark{Content: '', FG: body.Cyan},
	}

	VCS = vcs{
		Color: Seg[3],
		Split: "/",
		Join:  "│",
		Branch: vcsBranch{
			Empty:       body.Mark{Content: '○', FG: body.White},
			Initial:     body.Mark{Content: '', FG: body.LightGray},
			Tracked:     body.Mark{Content: '', FG: body.None},
			Detached:    body.Mark{Content: '', FG: body.Yellow},
			Ahead:       body.Mark{Content: '', FG: body.Green},
			Behind:      body.Mark{Content: '', FG: body.Blue},
			AheadBehind: body.Mark{Content: '', FG: body.Violet},
			Fatal:       body.Mark{Content: '', FG: body.Red},
			Unknown:     body.Mark{Content: '', FG: body.Cyan}},
		Status: vcsStatus{
			New:      body.Mark{Content: '', FG: body.LightGray},
			Added:    body.Mark{Content: '', FG: body.Green},
			Deleted:  body.Mark{Content: '', FG: body.Red},
			Modified: body.Mark{Content: '', FG: body.Yellow},
			Moved:    body.Mark{Content: '', FG: body.Blue},
			Copied:   body.Mark{Content: '', FG: body.Cyan},
			Unmerged: body.Mark{Content: '═', FG: body.Violet}},
	}
)
