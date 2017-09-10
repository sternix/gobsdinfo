package mem

import "github.com/sternix/gobsdinfo/util"

type InterruptStats struct {
	Version    string `json:"__version"`
	Statistics struct {
		Interrupts []struct {
			Name  string `json:"name"`
			Total int    `json:"total"`
			Rate  int    `json:"rate"`
		} `json:"interrupt"`
	} `json:"interrupt-statistics"`
	TotalInterrupts int `json:"total-interrupts"`
	TotalRate       int `json:"total-rate"`
}

func Interrupts() (InterruptStats, error) {
	var intr InterruptStats
	if err := util.ExecAndParse(&intr, "/usr/bin/vmstat", "-i"); err != nil {
		return intr, err
	}

	return intr, nil
}
