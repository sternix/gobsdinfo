package proc

import "github.com/sternix/gobsdinfo/util"

type Process struct {
	User          string `json:"user"`
	Pid           string `json:"pid"`
	CpuPercent    string `json:"percent-cpu"`
	MemoryPercent string `json:"percent-memory"`
	VirtualSize   string `json:"virtual-size"`
	Rss           string `json:"rss"`
	TerminalName  string `json:"terminal-name"`
	State         string `json:"state"`
	StartTime     string `json:"start-time"`
	CpuTime       string `json:"cpu-time"`
	Command       string `json:"command"`
}

type jProc struct {
	Information struct {
		Processes []Process `json:"process"`
	} `json:"process-information"`
}

func All() ([]Process, error) {
	var proc jProc
	if err := util.ExecAndParse(&proc, "/bin/ps", "-aux"); err != nil {
		return nil, err
	}
	return proc.Information.Processes, nil
}
