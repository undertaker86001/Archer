package server

type PostStartHookFunc func(context PostStartHookContext)

type PostStartHookContext struct {
	StopCh <-chan struct{}
}

type postStartHookEntry struct {
	hook PostStartHookFunc

	done chan struct{}
}

type PreShutdownHookContext struct {
	StopCh <-chan struct{}
}

type PreShutdownHookFunc func() error

type preShutdownHookEntry struct {
	hook PreShutdownHookFunc
}
