package sps

import (
	"time"

	"../ops"
)

var (
	conn = ops.Conn
	err  error
)

type spsModel struct {
	id       int
	name     string
	pn       string
	sn       string
	oldSim   string
	newSim   string
	position string
	status   string
	amount   int
	durables bool
	comment  string
	img1     string
	img2     string
	wanted   bool
	tags     []string
}

type spHistory struct {
	spid      int
	time      time.Time
	oldSim    string
	newSim    string
	oldPos    string
	newPos    string
	oldStatus string
	newStatus string
	user      string
	comment   string
	shorts    []string
}
