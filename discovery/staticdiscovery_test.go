package discovery

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
	yaml := []byte(`
nuvem.discovery.test.static.servers:
- localhost:8080
- 127.0.0.1:9080
`)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("nuvem.discovery.factory", "StaticDiscovery")

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//	servers := viper.GetStringSlice("nuvem.loadbalancer.test.static.servers")
	//	fmt.Printf("%+v\n", servers)
	//	factory := viper.GetString("nuvem.loadbalancer.test.factory")
	//	fmt.Printf("%+v\n", factory)

	discovery := Create("test")
	instances := assertDiscovery(t, discovery, 2)

	assertInstance(t, instances[0], "localhost", 8080)
	assertInstance(t, instances[1], "127.0.0.1", 9080)
}

func TestBuilder(t *testing.T) {
	println("\nTestBuilder")
	discovery := NewStaticBuilder("test").
		Servers("localhost:8080", "10.0.0.1:8765").
		Build()

	instances := assertDiscovery(t, discovery, 2)
	assertInstance(t, instances[0], "localhost", 8080)
	assertInstance(t, instances[1], "10.0.0.1", 8765)
}

func assertInstance(t *testing.T, instance util.Instance, host string, port int) {
	fmt.Printf("%+v\n", instance)

	assert.Equal(t, host, instance.Host, "wrong Host")
	assert.Equal(t, port, instance.Port, "wrong Port")
}

func assertDiscovery(t *testing.T, discovery Discovery, count int) []util.Instance {
	require.NotNil(t, discovery, "discovery was nil")

	instances := discovery.GetIntances()
	require.NotNil(t, instances, "instances was nil")
	assert.Equal(t, count, len(instances), "wrong # of instances")

	return instances
}
