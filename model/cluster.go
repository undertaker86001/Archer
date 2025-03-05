package model

import "fmt"

type (
	ClusterConfig struct {
		Name      string        `yaml:"name" json:"name"` // Name the cluster unique name
		TypeStr   string        `yaml:"type" json:"type"`
		Type      DiscoveryType `yaml:"discovery_type" json:"discovery_type"`
		Endpoints []*Endpoint   `yaml:"endpoints" json:"endpoints"`
		Lbstr     LbPolicyType  `yaml:"lbstr" json:"lbstr"`
	}

	//DisCoveryType
	DiscoveryType int32

	Endpoint struct {
		ID       string            `yaml:"ID" json:"ID"`                                                       // ID indicate one endpoint
		Name     string            `yaml:"name" json:"name"`                                                   // Name the cluster unique name
		Address  SocketAddress     `yaml:"socket_address" json:"socket_address" mapstructure:"socket_address"` // Address socket address
		Metadata map[string]string `yaml:"meta" json:"meta"`
	}

	SocketAddress struct {
		Address      string   `default:"0.0.0.0" yaml:"address" json:"address"`
		Port         string   `default:"8881" yaml:"port" json:"port"`
		ResolverName string   `yaml:"resolver_name" json:"resolver_name"`
		Domians      []string `yaml:"domians" json:"domians" mapstructure:"domains"`
	}
)

func (a SocketAddress) GetAddress() string {
	return fmt.Sprintf("%s:%v", a.Address, a.Port)
}
