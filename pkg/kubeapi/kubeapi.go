package kubeapi

import (
	"context"
	"github.com/rs/zerolog/log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// getKubeConfig check if the kubeconfig is set, if not use default
func getKubeConfig() string {
	env, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		temp := "~/.kube/config"
		return temp
	}
	return env
}

// NewKubeClient gets the k8s config and returns a clientset
func NewKubeClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Info().Msg("In-cluster config not available, looking for kubeconfig")
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", getKubeConfig())
		if err != nil {
			log.Error().Err(err).Msg("Failed to build config from kubeconfig")
			return nil, err
		}
	}
	client, err := kubernetes.NewForConfig(config)
	return client, err
}

// GetPodUIDList takes a clientset and returns a list of UIDs for pods based on a label selector and namespace
func GetPodUIDList(clientset kubernetes.Interface, namespace, labelSelector string) ([]string, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		log.Error().Err(err).Msgf("Failed to list pods in namespace %s with label selector %s", namespace, labelSelector)
		return nil, err
	}
	var podUIDs []string
	for _, pod := range pods.Items {
		podUIDs = append(podUIDs, string(pod.UID))
	}
	return podUIDs, nil
}
