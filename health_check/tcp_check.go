package health_check

import (
	"github.com/gobuffalo/logger"
	"net"
	"time"
)

// 通过tcp连接检测是否保活
type TCPChecker struct {
	timeout time.Duration
	addr    string
}

func (s *TCPChecker) CheckHealth() bool {
	conn, err := net.DialTimeout("tcp", s.addr, s.timeout)
	if err != nil {
		logger := logger.NewLogger("info")
		logger.Infof("[health check] tcp checker for host %s error: %v", s.addr, err)
		return false
	}
	conn.Close()
	return true
}

func (s *TCPChecker) OnTimeout() {}
