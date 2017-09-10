package mem

import "github.com/sternix/gobsdinfo/util"

type ForkStat struct {
	Version    string `json:"__version"`
	Statistics struct {
		Fork         int     `json:"fork"`
		ForkPages    int     `json:"fork-pages"`
		ForkAvarage  float32 `json:"fork-average"`
		VFork        int     `json:"vfork"`
		VForkPages   int     `json:"vfork-pages"`
		VForkAvarage float32 `json:"vfork-average"`
		RFork        int     `json:"rfork"`
		RForkPages   int     `json:"rfork-pages"`
		RForkAvarage float32 `json:"rfork-average"`
	} `json:"fork-statistics"`
}

func Forks() (ForkStat, error) {
	var fork ForkStat
	if err := util.ExecAndParse(&fork, "/usr/bin/vmstat", "-f"); err != nil {
		return fork, err
	}

	return fork, nil
}
