package fs

import "github.com/sternix/gobsdinfo/util"

type Filesystem struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	TotalBlocks       int    `json:"total-blocks"`
	UsedBlocks        int    `json:"used-blocks"`
	AvailableBlocks   int    `json:"available-blocks"`
	UsedPercent       int    `json:"used-percent"`
	InodesUsed        int    `json:"inodes-used"`
	InodesFree        int    `json:"inodes-free"`
	InodesUsedPercent int    `json:"inodes-used-percent"`
	MountedOn         string `json:"mounted-on"`
}

type storageInfo struct {
	Storage struct {
		Filesystems []Filesystem `json:"filesystem"`
	} `json:"storage-system-information"`
}

func Info() ([]Filesystem, error) {
	var sinfo storageInfo
	if err := util.ExecAndParse(&sinfo, "/bin/df", "-i", "-T"); err != nil {
		return nil, err
	}

	return sinfo.Storage.Filesystems, nil
}
