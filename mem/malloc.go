package mem

import "github.com/sternix/gobsdinfo/util"

type MallocStat struct {
	Version    string `json:"__version"`
	Statistics struct {
		Memories []struct {
			Type      string `json:"type"`
			InUse     int    `json:"in-use"`
			MemoryUse int    `json:"memory-use"`
			HighUse   string `json:"high-use"`
			Requests  int    `json:"requests"`
			Sizes     []int  `json:"size"`
		} `json:"memory"`
	} `json:"malloc-statistics"`
}

func Mallocs() (MallocStat, error) {
	var malloc MallocStat
	if err := util.ExecAndParse(&malloc, "/usr/bin/vmstat", "-m"); err != nil {
		return malloc, err
	}

	return malloc, nil
}
