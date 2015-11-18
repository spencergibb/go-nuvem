package serverlist

import (
	"bytes"
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactory(t *testing.T) {
	viper.SetConfigType("yaml")
	yaml := []byte(`
loadbalancer.test.serverlist.static.servers:
- localhost:8080
- 127.0.0.1:9080
`)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("loadbalancer.serverlist.test.factory", "StaticServerList")

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//	servers := viper.GetStringSlice("loadbalancer.test.static.servers")
	//	fmt.Printf("%+v\n", servers)
	//	factory := viper.GetString("loadbalancer.test.factory")
	//	fmt.Printf("%+v\n", factory)

	serverList := Create("test")
	servers := assertServerList(t, serverList, 2)

	assertServer(t, servers[0], "localhost", 8080)
	assertServer(t, servers[1], "127.0.0.1", 9080)
}

func TestBuilder(t *testing.T) {
	println("\nTestBuilder")
	serverList := NewStaticBuilder("test").
		Servers("localhost:8080", "10.0.0.1:8765").
		Build()

	servers := assertServerList(t, serverList, 2)
	assertServer(t, servers[0], "localhost", 8080)
	assertServer(t, servers[1], "10.0.0.1", 8765)
}

func assertServer(t *testing.T, server loadbalancer.Server, host string, port int) {
	fmt.Printf("%+v\n", server)

	assert.Equal(t, host, server.Host, "wrong Host")
	assert.Equal(t, port, server.Port, "wrong Port")
}

func assertServerList(t *testing.T, serverList ServerList, count int) []loadbalancer.Server {
	require.NotNil(t, serverList, "serverList was nil")

	servers := serverList.GetServers()
	require.NotNil(t, servers, "servers was nil")
	assert.Equal(t, count, len(servers), "wrong # of servers")

	return servers
}
