package model

type Bootstrap struct {
	StaticResources
}

type StaticResources struct {
	Clusters []*ClusterConfig `yaml:"clusters" json:"clusters" mapstructure:"clusters"`
}

type ShutdownConfig struct {
	Timeout      string `default:"60s" yaml:"timeout" json:"timeout,omitempty"`
	StepTimeout  string `default:"10s" yaml:"step_timeout" json:"timeout,omitempty"`
	RejectPolicy string `default:"immediacy" yaml:"reject_policy" json:"reject_policy,omitempty"`
}
