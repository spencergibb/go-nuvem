package serverlist

import (
	"bytes"
	"fmt"
	"github.com/spencergibb/go-nuvem/discovery"
	"github.com/spf13/viper"
	"testing"
)

func TestDiscoveryFactory(t *testing.T) {
	viper.SetConfigType("yaml")
	yaml := []byte(`
nuvem.discovery.testdsl.static.servers:
- localhost:8080
- 127.0.0.1:9080
`)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("nuvem.discovery.testdsl.factory", discovery.StaticFactoryKey)
	viper.SetDefault("nuvem.loadbalancer.serverlist.testdsl.factory", DisoveryFactoryKey)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//	servers := viper.GetStringSlice("nuvem.loadbalancer.test.static.servers")
	//	fmt.Printf("%+v\n", servers)
	//	factory := viper.GetString("nuvem.loadbalancer.test.factory")
	//	fmt.Printf("%+v\n", factory)

	serverList := Create("testdsl")
	servers := assertServerList(t, serverList, 2)

	assertServer(t, servers[0], "localhost", 8080)
	assertServer(t, servers[1], "127.0.0.1", 9080)
}

func TestDiscoveryBuilder(t *testing.T) {
	println("\nTestBuilder")
	serverList := NewDiscoveryBuilder("testdsl").
		Discovery(discovery.NewStaticBuilder("testdsl").Servers("localhost:8080", "10.0.0.1:8765").Build()).
		Build()

	servers := assertServerList(t, serverList, 2)
	assertServer(t, servers[0], "localhost", 8080)
	assertServer(t, servers[1], "10.0.0.1", 8765)
}
