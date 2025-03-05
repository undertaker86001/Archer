package model

type LbPolicyType string

const (
	LoadBalancerRoundRobin LbPolicyType = "RoundRobin"
)

var LbPolicyTypeValue = map[string]LbPolicyType{
	"RoundRobin": LoadBalancerRoundRobin,
}
