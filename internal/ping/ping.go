package ping

import (
	probing "github.com/prometheus-community/pro-bing"
)

func Ping(host string) bool {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		return false
	}
	pinger.Count = 1
	err = pinger.Run()
	if err != nil {
		return false
	}

	stats := pinger.Statistics()
	return stats.PacketsRecv > 0
}
