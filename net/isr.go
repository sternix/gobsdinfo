package net

import "github.com/sternix/gobsdinfo/util"

type Isr struct {
	Protocols []struct {
		Name       string `json:"name"`
		Protocol   int    `json:"protocol"`
		QueueLimit int    `json:"queue-limit"`
		PolicyType string `json:"policy-type"`
		Policy     string `json:"policy"`
		Flags      string `json:"flags"`
	} `json:"protocol"`
	Workstream []struct {
		Work []struct {
			Workstream       int    `json:"workstream"`
			Cput             int    `json:"cpu"`
			Name             string `json:"name"`
			Length           int    `json:"length"`
			Watermark        int    `json:"watermark"`
			Dispatched       int    `json:"dispatched"`
			HybridDispatched int    `json:"hybrid-dispatched"`
			QueueDrops       int    `json:"queue-drops"`
			Queued           int    `json:"queued"`
			Handled          int    `json:"handled"`
		} `json:"work"`
	} `json:"workstream"`
}

type jIsr struct {
	NetIsr struct {
		isr Isr
	} `json:"netisr"`
}

func Isrs() (Isr, error) {
	var isr jIsr
	if err := util.ExecAndParse(&isr, "/usr/bin/netstat", "-Q"); err != nil {
		return isr.NetIsr.isr, err
	}
	return isr.NetIsr.isr, nil
}
