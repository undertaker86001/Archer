package config

import (
	"cluster_manager/health_check"
	"cluster_manager/logger"
	"cluster_manager/server"
	restful "github.com/emicklei/go-restful"
	"net/http"
	"sync"
	"time"
)

type Config struct {
	ExternalAddress       string
	HandlerChainWaitGroup *sync.WaitGroup
	HealthChecks          health_check.HealthChecker
	BuildHandlerChainFunc func(apiHandler http.Handler, c *Config) http.Handler
	RequestTimeout        time.Duration
}

type CompletedConfig struct {
	*Config
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

type HandlerChainBuilderFn func(apiHandler http.Handler) http.Handler

func (c CompletedConfig) New(name string) (*server.GenericServer, error) {
	handlerChainBuilder := func(handler http.Handler) http.Handler {
		return c.BuildHandlerChainFunc(handler, c.Config)
	}

	handler := NewServerHandler(name, handlerChainBuilder, nil)

	if handler != nil {
		logger.Info("[error] handler not extists")
	}
	return &server.GenericServer{}, nil
}

func NewServerHandler(name string, handlerChainBuilder HandlerChainBuilderFn, notFoundHandler http.Handler) *http.Handler {

	gorestfulContainer := restful.NewContainer()
	gorestfulContainer.ServeMux = http.NewServeMux()
	gorestfulContainer.Router(restful.CurlyRouter{})

	director := server.director()
	return nil
}
