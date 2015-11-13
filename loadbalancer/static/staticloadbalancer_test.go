package static

import (
	"testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"fmt"
	"bytes"
	"github.com/spencergibb/go-nuvem/loadbalancer"
)

func TestChoose(t *testing.T) {

	loadbalancer.New = NewStaticLoadBalancer

	viper.SetConfigType("yaml")
	yaml := []byte(`
loadbalancer.static.test.servers:
- localhost:8080
`)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
//	servers := viper.GetStringSlice("loadbalancer.static.servers")
//	fmt.Printf("%+v\n", servers)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	lb := loadbalancer.New("test")
	server := lb.Choose()
	assert.NotNil(t, server, "server was nil")
	fmt.Printf("%+v\n", server)

	assert.Equal(t, server.Host, "localhost", "wrong Host")
	assert.Equal(t, server.Port, 8080, "wrong Port")
}