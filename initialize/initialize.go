package initialze

import "github.com/spencergibb/go-nuvem/loadbalancer/noop"

func Init() {
	noop.Load()
}
