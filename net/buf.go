package net

import "github.com/sternix/gobsdinfo/util"

type Mbuf struct {
	MbufCurrent                int `json:"mbuf-current"`
	MbufCache                  int `json:"mbuf-cache"`
	MbufTotal                  int `json:"mbuf-total"`
	ClusterCurrent             int `json:"cluster-current"`
	ClusterCache               int `json:"cluster-cache"`
	ClusterTotal               int `json:"cluster-total"`
	ClusterMax                 int `json:"cluster-max"`
	PacketCount                int `json:"packet-count"`
	PacketFree                 int `json:"packet-free"`
	JumboCount                 int `json:"jumbo-count"`
	JumboCache                 int `json:"jumbo-cache"`
	JumboTotal                 int `json:"jumbo-total"`
	JumboMax                   int `json:"jumbo-max"`
	JumboPageSize              int `json:"jumbo-page-size"`
	Jumbo9Count                int `json:"jumbo9-count"`
	Jumbo9Cache                int `json:"jumbo9-cache"`
	Jumbo9Total                int `json:"jumbo9-total"`
	Jumbo9Max                  int `json:"jumbo9-max"`
	Jumbo16Count               int `json:"jumbo16-count"`
	Jumbo16Cache               int `json:"jumbo16-cache"`
	Jumbo16Total               int `json:"jumbo16-total"`
	Jumbo16Limit               int `json:"jumbo16-limit"`
	BytesInUse                 int `json:"bytes-in-use"`
	BytesInCache               int `json:"bytes-in-cache"`
	BytesTotal                 int `json:"bytes-total"`
	MbufFailures               int `json:"mbuf-failures"`
	ClusterFailures            int `json:"cluster-failures"`
	PacketFailures             int `json:"packet-failures"`
	MbufSleeps                 int `json:"mbuf-sleeps"`
	ClusterSleeps              int `json:"cluster-sleeps"`
	PacketSleeps               int `json:"packet-sleeps"`
	JumbopSleeps               int `json:"jumbop-sleeps"`
	Jumbo9Sleeps               int `json:"jumbo9-sleeps"`
	Jumbo16Sleeps              int `json:"jumbo16-sleeps"`
	JumbopFailures             int `json:"jumbop-failures"`
	Jumbo9Failures             int `json:"jumbo9-failures"`
	Jumbo16Failures            int `json:"jumbo16-failures"`
	SendfileSyscalls           int `json:"sendfile-syscalls"`
	SendfileNoIo               int `json:"sendfile-no-io"`
	SendfileIoCount            int `json:"sendfile-io-count"`
	SendfilePagesSent          int `json:"sendfile-pages-sent"`
	SendfilePagesValid         int `json:"sendfile-pages-valid"`
	SendfileRequestedReadahead int `json:"sendfile-requested-readahead"`
	SendfileReadahead          int `json:"sendfile-readahead"`
	SendfileBusyEncounters     int `json:"sendfile-busy-encounters"`
	SfbufsAllocFailed          int `json:"sfbufs-alloc-failed"`
	SfbufsAllocWait            int `json:"sfbufs-alloc-wait"`
}

type jMbuf struct {
	Stat struct {
		mbuf Mbuf
	} `json:"mbuf-statistics"`
}

func Mbufs() (Mbuf, error) {
	var mbuf jMbuf
	if err := util.ExecAndParse(&mbuf, "/usr/bin/netstat", "-m"); err != nil {
		return mbuf.Stat.mbuf, err
	}
	return mbuf.Stat.mbuf, nil
}
