package health_check

// 健康检查配置项
type HealthCheckConfig struct {
	Protocol            string                 `json:"protocol,omitempty"`
	TimeoutConfig       string                 `json:"timeout,omitempty"`
	IntervalConfig      string                 `json:"interval,omitempty"`
	InitialDelaySeconds string                 `json:"initial_delay_seconds,omitempty"`
	SessionConfig       map[string]interface{} `json:"check_config,omitempty"`
	HealthyThreshold    uint32                 `json:"healthy_threshold,omitempty"`
}
