package net

import "github.com/sternix/gobsdinfo/util"

type InetSocket struct {
	Protocol            string `json:"protocol"`
	ReveiveBytesWaiting int    `json:"receive-bytes-waiting"`
	SendBytesWaiting    int    `json:"send-bytes-waiting"`
	Local               struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"local"`
	Remote struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"remote"`
}

type jInetSocketStat struct {
	Stat struct {
		Sockets []InetSocket `json:"socket"`
	} `json:"statistics"`
}

func InetSocketStats() ([]InetSocket, error) {
	return inetSocketStats("inet")
}

func Inet6SocketStats() ([]InetSocket, error) {
	return inetSocketStats("inet6")
}

func inetSocketStats(af string) ([]InetSocket, error) {
	var sstat jInetSocketStat

	if err := util.ExecAndParse(&sstat, "/usr/bin/netstat", "-a", "-n", "-f", af); err != nil {
		return nil, err
	}
	return sstat.Stat.Sockets, nil
}
