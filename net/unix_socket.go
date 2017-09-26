package net

import "github.com/sternix/gobsdinfo/util"

type UnixSocket struct {
	Address             string `json:"address"`
	Type                string `json:"type"`
	ReceiveBytesWaiting int    `json:"receive-bytes-waiting"`
	SendBytesWaiting    int    `json:"send-bytes-waiting"`
	Vnode               string `json:"vnode"`
	Connection          string `json:"connection"`
	FirstReference      string `json:"first-reference"`
	NextReference       string `json:"next-reference"`
	Path                string `json:"path,omitempty"`
}

type jUnixSocketStat struct {
	Stat struct {
		Sockets []UnixSocket `json:"socket"`
	} `json:"statistics"`
}

func UnixSocketStats() ([]UnixSocket, error) {
	var sstat jUnixSocketStat

	if err := util.ExecAndParse(&sstat, "/usr/bin/netstat", "-a", "-n", "-f", "unix"); err != nil {
		return nil, err
	}
	return sstat.Stat.Sockets, nil
}
