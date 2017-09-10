package mem

import "github.com/sternix/gobsdinfo/util"

type ObjectStat struct {
	Version string `json:"__version"`
	Objects []struct {
		Resident    int    `json:"resident"`
		Active      int    `json:"active"`
		Inactive    int    `json:"inactive"`
		RefCount    int    `json:"refcount"`
		ShadowCount int    `json:"shadowcount"`
		Attribute   string `json:"attirbute"`
		Type        string `json:"type"`
		Path        string `json:"path"`
	} `json:"object"`
}

func Objects() (ObjectStat, error) {
	var obj ObjectStat
	if err := util.ExecAndParse(&obj, "/usr/bin/vmstat", "-o"); err != nil {
		return obj, err
	}

	return obj, nil
}
