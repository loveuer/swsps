package sps

import (
	"time"

	"../ops"
)

var (
	conn = ops.Conn
	err  error
)

// SpsModel ...
type SpsModel struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	PN           string   `json:"pn"`
	SN           string   `json:"sn"`
	OrgSim       string   `json:"orgsim"`
	NowSim       string   `json:"nowsim"`
	Position     string   `json:"pos"`
	Status       string   `json:"status"`
	Amount       int      `json:"amount"`
	Limit        int      `json:"limit"`
	IsConsumable bool     `json:"is_consumable"`
	Comment      string   `json:"comment"`
	Img1         string   `json:"img1"`
	Img2         string   `json:"img2"`
	Tags         []string `json:"tags"`
	// wanted       bool
}

// SpHistory ...
type SpHistory struct {
	Spid       int       `json:"spid"`
	Time       time.Time `json:"time"`
	SD         string    `json:"simple_date"`
	ST         string    `json:"simple_time"`
	LastSim    string    `json:"last_sim"`
	NextSim    string    `json:"next_sim"`
	LastPos    string    `json:"last_pos"`
	NextPos    string    `json:"next_pos"`
	LastStatus string    `json:"last_sts"`
	NextStatus string    `json:"next_sts"`
	AmountChg  int       `json:"amount_chg"`
	Auth       string    `json:"auth"`
	Comment    string    `json:"comment"`
	Short      []string  `json:"short"`
}
