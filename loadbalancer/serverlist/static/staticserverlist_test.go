package static

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist/builder"
	"github.com/spencergibb/go-nuvem/initialize"
)

func TestChoose(t *testing.T) {
	initialze.Init() //TODO: get rid of
	viper.SetConfigType("yaml")
	yaml := []byte(`
loadbalancer.test.serverlist.static.servers:
- localhost:8080
- 127.0.0.1:9080
`)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("loadbalancer.serverlist.test.factory", "StaticServerList")
//	servers := viper.GetStringSlice("loadbalancer.test.static.servers")
//	fmt.Printf("%+v\n", servers)
//	factory := viper.GetString("loadbalancer.test.factory")
//	fmt.Printf("%+v\n", factory)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	sl := builder.Build("test")
	require.NotNil(t, sl, "sl was nil")

	servers := sl.GetServers()
	require.NotNil(t, servers, "servers was nil")
	assert.Equal(t, 2, len(servers), "wrong # of servers")

	server := servers[0]
	fmt.Printf("%+v\n", server)

	assert.Equal(t, "localhost", server.Host, "wrong Host")
	assert.Equal(t, 8080, server.Port, "wrong Port")
}