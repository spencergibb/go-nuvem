package serverlist

func NewStaticBuilder(ns string) StaticBuilder {
	b := StaticBuilder{namespace: ns}
	return b
}

type StaticBuilder struct {
	namespace string
	servers   []string
}

func (b StaticBuilder) Servers(servers ...string) StaticBuilder {
	b.servers = servers
	return b
}

func (b StaticBuilder) Build() ServerList {
	serverList := StaticServerList{}
	serverList.Namespace = b.namespace
	serverList.Servers = b.servers
	return &serverList
}
