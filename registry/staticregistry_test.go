package registry

import (
	"bytes"
	"fmt"
	"github.com/spencergibb/go-nuvem/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactory(t *testing.T) {
	viper.SetConfigType("yaml")
	yaml := []byte(``)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("nuvem.registry.factory", StaticFactoryKey)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//	servers := viper.GetStringSlice("nuvem.loadbalancer.test.static.servers")
	//	fmt.Printf("%+v\n", servers)
	//	factory := viper.GetString("nuvem.loadbalancer.test.factory")
	//	fmt.Printf("%+v\n", factory)

	registry := Create("testreg")

	servers := assertRegistry(t, registry)
	assertServers(t, servers)
}

func TestBuilder(t *testing.T) {
	println("\nTestBuilder")
	registry := NewStaticBuilder("testreg").Build()

	servers := assertRegistry(t, registry)
	assertServers(t, servers)
}

func assertServers(t *testing.T, servers []string) {
	set := make(map[string]bool)
	for _, v := range servers {
		set[v] = true
	}

	assert.True(t, set["127.0.0.1:8080"], "missing 127...:8080")
	assert.True(t, set["localhost:9080"], "missing local...:9080")
}

func assertRegistry(t *testing.T, registry Registry) []string {
	registry.Register(util.Instance{Id: "1", Host: "127.0.0.1", Port: 8080})
	registry.Register(util.Instance{Id: "2", Host: "localhost", Port: 9080})

	servers := viper.GetStringSlice(util.GetStaticRegistryKey(registry))

	require.NotNil(t, servers, "servers was nil")
	require.Equal(t, 2, len(servers), "wrong # of servers")

	return servers
}
