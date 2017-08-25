package kubernetes

var (
	apiHost, cpuLimit, cpuRequest, memoryLimit, memoryRequest, namespace string
	replicas                                                             int
	EnableKubernetes                                                     bool
)

func init() {
	flag.StringVar(&apiHost, "api-host", "127.0.0.1:8001", "Kubernetes API server")
	flag.StringVar(&cpuLimit, "cpu-limit", "100m", "Max CPU in milicores")
	flag.StringVar(&cpuRequest, "cpu-request", "100m", "Min CPU in milicores")
	flag.StringVar(&memoryLimit, "memory-limit", "64M", "Max memory in MB")
	flag.StringVar(&memoryRequest, "memory-request", "64M", "Min memory in MB")
	flag.StringVar(&namespace, "namespace", "default", "The Kubernetes namespace.")
	flag.IntVar(&replicas, "replicas", 1, "Number of replicas")
	flag.BoolVar(&EnableKubernetes, "kubernetes", false, "Deploy to Kubernetes.")
}
