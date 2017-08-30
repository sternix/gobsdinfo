package host

import "github.com/sternix/gobsdinfo/util"

type LoggedUser struct {
	Name      string `json:"user"`
	Tty       string `json:"tty"`
	From      string `json:"from"`
	LoginTime string `json:"login-time"`
	Command   string `json:"command"`
}

type Uptime struct {
	TimeOfDay     string `json:"time-of-day"`
	Uptime        int64  `json:"uptime"`
	Days          int    `json:"days"`
	Hours         int    `json:"hours"`
	Minutes       int    `json:"minutes"`
	Seconds       int    `json:"seconds"`
	UptimeHuman   string `json:"uptime-human"`
	Users         int    `json:"users"`
	LoadAvarage1  string `json:"load-avarage-1"`
	LoadAvarage5  string `json:"load-avarage-5"`
	LoadAvarage15 string `json:"load-avarage-15"`
	LoggedUsers   struct {
		Users []LoggedUser `json:"user-entry"`
	} `json:"user-table"`
}

type jUptime struct {
	Uptime Uptime `json:"uptime-information"`
}

func UptimeInfo() (*Uptime, error) {
	var uptime jUptime
	if err := util.ExecAndParse(&uptime, "/usr/bin/w"); err != nil {
		return nil, err
	}

	return &uptime.Uptime, nil
}
