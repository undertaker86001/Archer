package resource

type ComputeOptions struct {
	ContainerNum   int
	RuntimeCmd     string
	RunnerSpecPath string
	RunnerDataPath string
	ConfPath       string
	RuntimePath    string
	CachePath      string
	InvokerSocks   string
	ListenPath     string
	ListenType     string
}
