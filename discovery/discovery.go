package discovery

type Instance interface {
	GetId() string
	GetHost() string
	GetPort() int
}

type Discovery interface {
	GetIntances() []Instance
}
