package arp

import "github.com/sternix/gobsdinfo/util"

type Entry struct {
	Hostname   string `json:"hostname"`
	IpAddress  string `json:"ip-address"`
	MacAddress string `json:"mac-address"`
	Interface  string `json:"interface"`
	Expires    int    `json:"expires,omitempty"`
	Permanent  bool   `json:"permanent,omitempty"`
	Type       string `json:"type"`
}

type jArp struct {
	Version string `json:"__version"`
	Arp     struct {
		Entries []Entry `json:"arp-cache"`
	} `json:"arp"`
}

func Entries() ([]Entry, error) {
	var arp jArp
	if err := util.ExecAndParse(&arp, "/usr/sbin/arp", "-a"); err != nil {
		return nil, err
	}
	return arp.Arp.Entries, nil
}
