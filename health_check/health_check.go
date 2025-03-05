package health_check

import (
	"cluster_manager/logger"
	"cluster_manager/model"
	logger2 "github.com/gobuffalo/logger"
	"time"
)

const (
	DefaultTimeout                 = time.Second
	DefaultInterval                = 3 * time.Second
	DefaultHealthyThreshold uint32 = 5
)

// 通过网络监听器配置的健康检查配置
type HealthChecker struct {
	checkers      map[string]*EndpointChecker
	sessionConfig map[string]interface{}

	cluster        *model.ClusterConfig
	timeout        time.Duration //tcp连接的超时时间
	intervalBase   time.Duration
	healthThresold time.Duration
}

// worker端点的单健康检查配置
type EndpointChecker struct {
	endpoint      *model.Endpoint
	HealthChecker *HealthChecker
	tcpChecker    *TCPChecker
}

func (hc *HealthChecker) StopOne(endpoint *model.Endpoint) {
	hc.StopOne(endpoint)
}

func (hc *HealthChecker) StartOne(endpoint *model.Endpoint) {

}

func (hc *EndpointChecker) Start() {

}

func newChecker(endpoint *model.Endpoint, hc *HealthChecker) *EndpointChecker {
	c := &EndpointChecker{
		endpoint: endpoint,
	}
	return c
}

func (hc *HealthChecker) startCheck(endpoint *model.Endpoint) {
	addr := endpoint.Address.GetAddress()
	if _, ok := hc.checkers[addr]; !ok {
		c := newChecker(endpoint, hc)
		hc.checkers[addr] = c
		go c.Start()
		logger := logger2.NewLogger("info")
		logger.Infof("[health check] create a health check session for %s", addr)
	}
}

func CreateHealthCheck(cluster *model.ClusterConfig, cfg HealthCheckConfig) *HealthChecker {
	timeout, err := time.ParseDuration(cfg.TimeoutConfig)
	if err != nil {
		timeout = DefaultTimeout
		logger.Debug("[health check] 超时时间解析出错 %s", err)
	}

	interval, err := time.ParseDuration(cfg.IntervalConfig)
	if err != nil {
		interval = DefaultInterval
		logger.Errorf("[health check] 超时间隔解析出错 $s", err)
	}

	healthThreshold := cfg.HealthyThreshold
	if healthThreshold == 0 {
		healthThreshold = DefaultHealthyThreshold
		logger.Errorf("[health check] 健康检查周期出错 $s", err)
	}

	return &HealthChecker{
		cluster:        cluster,
		sessionConfig:  cfg.SessionConfig,
		timeout:        timeout,
		intervalBase:   interval,
		healthThresold: time.Duration(healthThreshold),
	}
}
