package discovery

func NewStaticBuilder(ns string) StaticBuilder {
	b := StaticBuilder{namespace: ns}
	return b
}

type StaticBuilder struct {
	namespace string
	instances []string
}

func (b StaticBuilder) Servers(instances ...string) StaticBuilder {
	b.instances = instances
	return b
}

func (b StaticBuilder) Build() Discovery {
	discovery := StaticDiscovery{}
	discovery.Namespace = b.namespace
	discovery.Instances = b.instances
	return &discovery
}
