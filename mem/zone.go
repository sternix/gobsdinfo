package mem

import "github.com/sternix/gobsdinfo/util"

type Zone struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Limit    int    `json:"limit"`
	Used     int    `json:"used"`
	Free     int    `json:"free"`
	Requests int    `json:"requests"`
	Fail     int    `json:"fail"`
	Sleep    int    `json:"sleep"`
}

type jMemZone struct {
	Version    string `json:"__version"`
	Statistics struct {
		Zones []Zone `json:"zone"`
	} `json:"memory-zone-statistics"`
}

func Zones() ([]Zone, error) {
	var zone jMemZone
	if err := util.ExecAndParse(&zone, "/usr/bin/vmstat", "-z"); err != nil {
		return nil, err
	}

	return zone.Statistics.Zones, nil
}
