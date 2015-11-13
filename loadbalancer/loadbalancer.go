package loadbalancer

type LoadBalancer interface {

	choose() *Server
}
