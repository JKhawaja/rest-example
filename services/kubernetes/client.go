package kubernetes

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	apiHost, namespace                       string
	cpuMax, cpuMin, memMax, memMin, replicas int
	// Deploy allows one to specify in a seperate package whether or not the app should be deployed to kubernetes
	Deploy bool
)

// TODO: default values should be specifiable by environment variables (or a config file)(??)
func init() {
	flag.StringVar(&apiHost, "api", "127.0.0.1:8001", "K8s API server")
	flag.IntVar(&cpuMax, "cpumax", 100, "Maximum CPU (in milicores)")
	flag.IntVar(&cpuMin, "cpumin", 100, "Minimum CPU (in milicores)")
	flag.IntVar(&memMax, "memmax", 64, "Maximum memory (in MBs)")
	flag.IntVar(&memMin, "memmin", 64, "Minimum memory (in MBs)")
	flag.StringVar(&namespace, "namespace", "default", "Namespace in K8s cluster.")
	flag.IntVar(&replicas, "replicas", 1, "Replica quantity.")
	flag.BoolVar(&Deploy, "kubernetes", false, "Specifies to deploy to Kubernetes (or not).")
}

// DeploymentConfig are the configuration parameters for a deployment (constructed from the flags)
type DeploymentConfig struct {
	Annotations map[string]string
	Args        []string
	Env         map[string]string
	BURL        string // binary url
	Name        string
	Labels      map[string]string
}

// K8Client ...
type K8Client struct {
	ClientSet *kubernetes.Clientset
}

// NewKubernetesClient ...
func NewKubernetesClient(path string) (Client, error) {
	emptyClient := &K8Client{}

	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return emptyClient, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return emptyClient, err
	}

	return &K8Client{
		ClientSet: clientset,
	}, nil
}

// CreateDeployment ...
func (c *K8Client) CreateDeployment(config DeploymentConfig) error {
	// create deployment client
	deployClient := c.ClientSet.AppsV1beta1().Deployments(apiv1.NamespaceDefault)

	// define resource requirements
	resources := apiv1.ResourceRequirements{
		Limits:   make(apiv1.ResourceList),
		Requests: make(apiv1.ResourceList),
	}

	// resource maximums (for container)
	resources.Limits["cpu"] = *resource.NewQuantity(int64(cpuMax), resource.DecimalSI)
	resources.Limits["memory"] = *resource.NewQuantity(int64(memMax), resource.DecimalSI)

	// resource minimums (for container)
	resources.Requests["cpu"] = *resource.NewQuantity(int64(cpuMin), resource.DecimalSI)
	resources.Requests["memory"] = *resource.NewQuantity(int64(memMin), resource.DecimalSI)

	// path to binary (on the volume that will be mounted to the container)
	binPath := filepath.Join("/opt/bin", config.Name)

	// app container
	container := apiv1.Container{
		Args:            config.Args,
		Command:         []string{binPath},
		Image:           "gliderlabs/alpine",
		ImagePullPolicy: "Always",
		Name:            config.Name,
		VolumeMounts: []apiv1.VolumeMount{
			{
				Name:      "bin",
				MountPath: "/opt/bin",
			},
		},
		Resources: resources,
	}

	// create environment variables for app container
	if len(config.Env) > 0 {
		var env []apiv1.EnvVar
		for name, value := range config.Env {
			env = append(env, apiv1.EnvVar{Name: name, Value: value})
		}
		container.Env = env
	}

	// initi containers for Pod
	initContainers := []apiv1.Container{
		// First container: downloads binary to Volume
		{
			Name:            "download",
			Image:           "gliderlabs/alpine",
			ImagePullPolicy: "Always",
			Command:         []string{"wget", "-O", binPath, config.BURL},
			VolumeMounts: []apiv1.VolumeMount{
				{
					Name:      "bin",
					MountPath: "/opt/bin",
				},
			},
		},
		// Second container: ensures that binary on volume is executable
		{
			Name:            "chmod",
			Image:           "gliderlabs/alpine",
			ImagePullPolicy: "Always",
			Command:         []string{"chmod", "+x", binPath},
			VolumeMounts: []apiv1.VolumeMount{
				{
					Name:      "bin",
					MountPath: "/opt/bin",
				},
			},
		},
	}

	// Label: "run":"app_name"
	config.Labels["run"] = config.Name

	deployment := &v1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      config.Name,
			Namespace: namespace,
		},
		Spec: v1beta1.DeploymentSpec{
			Replicas: int32Ptr(int32(replicas)),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"run": config.Name,
					},
					Annotations: config.Annotations,
				},
				Spec: apiv1.PodSpec{
					Containers:     []apiv1.Container{container},
					InitContainers: initContainers,
					Volumes: []apiv1.Volume{
						{
							Name:         "bin",
							VolumeSource: apiv1.VolumeSource{},
						},
					},
				},
			},
		},
	}

	// Create deployment
	log.Println("Creating Deployment ...")
	result, err := deployClient.Create(deployment)
	if err != nil {
		return err
	}

	log.Printf("Pod %s has deployed on cluster %s in namespace %s", result.GetName(), result.GetClusterName(), result.GetNamespace())

	return nil
}

func int32Ptr(i int32) *int32 { return &i }
