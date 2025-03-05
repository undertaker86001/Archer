package server

import (
	"sync"
	"time"
)

type GenericServer struct {
	ShutdownTimeout time.Duration

	postStartHookLock sync.Mutex
	postStartHooks    map[string]PostStartHookFunc
}
