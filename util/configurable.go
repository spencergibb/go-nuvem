package util

type Configurable interface {
	Configure(namespace string)
	GetNamespace() string
}
