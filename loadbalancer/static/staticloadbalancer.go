package static

import (
    "github.com/spencergibb/go-nuvem/loadbalancer"
    "github.com/spf13/viper"
	"math/rand"
	"net"
	"strconv"
//	"fmt"
)

type StaticLoadBalancer struct {

}

func (s StaticLoadBalancer) choose() *loadbalancer.Server {
	servers := viper.GetStringSlice("loadbalancer.static.servers")

	if (len(servers) == 0) {
		return nil
	}

//	fmt.Printf("%+v\n", servers)
//	TODO: implement rules
	idx := rand.Intn(len(servers))

	host, portStr, err := net.SplitHostPort(servers[idx])

	port, err := strconv.Atoi(portStr)

	print(err)

	return &loadbalancer.Server{Host: host, Port: port}
}