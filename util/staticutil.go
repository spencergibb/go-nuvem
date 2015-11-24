package util

import "fmt"

func GetStaticRegistryKey(configurable Configurable) string {
	//TODO: change key, remove discovery
	key := fmt.Sprintf("nuvem.discovery.%s.static.servers", configurable.GetNamespace())
	fmt.Printf("key %+v\n", key)
	return key
}
