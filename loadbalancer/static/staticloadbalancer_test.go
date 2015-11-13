package static

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	lbbuilder "github.com/spencergibb/go-nuvem/loadbalancer/builder"
	"github.com/spencergibb/go-nuvem/initialize"
)

func TestChoose(t *testing.T) {
	initialze.Init() //TODO: get rid of
	viper.SetConfigType("yaml")
	yaml := []byte(`
loadbalancer.test.static.servers:
- localhost:8080
`)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("loadbalancer.test.factory", "StaticLoadBalancer")
//	servers := viper.GetStringSlice("loadbalancer.test.static.servers")
//	fmt.Printf("%+v\n", servers)
//	factory := viper.GetString("loadbalancer.test.factory")
//	fmt.Printf("%+v\n", factory)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

//	loadbalancer.New = NewStaticLoadBalancer
//	lb := loadbalancer.New()
	lb := lbbuilder.Build("test")
	require.NotNil(t, lb, "lb was nil")

	server := lb.Choose()
	require.NotNil(t, server, "server was nil")
	fmt.Printf("%+v\n", server)

	assert.Equal(t, "localhost", server.Host, "wrong Host")
	assert.Equal(t, 8080, server.Port, "wrong Port")
}