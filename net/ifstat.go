package net

import "github.com/sternix/gobsdinfo/util"

type IfStat struct {
	Name             string `json:"name"`
	Flags            string `json:"flags"`
	MTU              int    `json:"mtu,omitempty"`
	Network          string `json:"network"`
	Address          string `json:"address"`
	ReceivedPackets  uint64 `json:"received-packets"`
	ReceivedErrors   uint64 `json:"received-errors,omitempty"`
	InDroppedPackets uint64 `json:"dropped-packets,omitempty"`
	RecivedBytes     uint64 `json:"received-bytes"`
	SentPackets      uint64 `json:"sent-packets"`
	SendErrors       uint64 `json:"send-errors,omitempty"`
	SentBytes        uint64 `json:"sent-bytes"`
	Collisions       uint64 `json:"collisions,omitempty"`
	// Bug duplicate key dropped-packets, reported FreeBSD developer via email
	//OutDroppedPackets uint64 `json:"dropped-packets,omitempty"`
}

type jIfStat struct {
	Stats struct {
		IfStats []IfStat `json:"interface"`
	} `json:"statistics"`
}

func IfStats() ([]IfStat, error) {
	var ifstat jIfStat
	if err := util.ExecAndParse(&ifstat, "/usr/bin/netstat", "-i", "-b", "-d", "-n", "-W"); err != nil {
		return nil, err
	}
	return ifstat.Stats.IfStats, nil
}
