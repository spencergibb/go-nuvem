package noop

import (
	"bytes"
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/builder"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChoose(t *testing.T) {
	viper.SetConfigType("yaml")
	yaml := []byte(``)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("loadbalancer.test.factory", "NoopLoadBalancer")

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	lb := builder.Build("test")
	assertLoadBalancer(t, lb)
}

func TestBuilder(t *testing.T) {
	println("\nTestBuilder")
	b := NewBuilder("test")
	lb := b.Build()

	assertLoadBalancer(t, lb)
}

func assertLoadBalancer(t *testing.T, lb loadbalancer.LoadBalancer) {
	require.NotNil(t, lb, "lb was nil")

	server := lb.Choose()
	assert.Nil(t, server, "server was not nil")
}
