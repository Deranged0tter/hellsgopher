package hellsgopher

import (
	"github.com/go-ping/ping"
)

// ping a given ip
// returns true if host is up
func Ping(ip string) (bool, error) {
	var isUp bool

	// new pinger
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return false, err
	}

	pinger.Count = 3

	pinger.OnFinish = func(s *ping.Statistics) {
		if s.PacketsRecv == 3 {
			isUp = true
		} else {
			isUp = false
		}
	}

	err = pinger.Run()
	if err != nil {
		return false, err
	}

	return isUp, nil
}
