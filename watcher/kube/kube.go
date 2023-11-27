package kube

import (
	"log"
	"sync"

	"github.com/pigeoncorp/bosun/watcher/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset *kubernetes.Clientset
var clientLock = &sync.Mutex{}

// GetClientConfig returns a cached REST client config if available or attempts to fetch one from kubernetes
func getClientset() *kubernetes.Clientset {
	clientLock.Lock()
	defer clientLock.Unlock()

	if clientset == nil {
		cc, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalf("unable to obtain in-cluster client configuration\n%v", err)
		}

		cs, err := kubernetes.NewForConfig(cc)
		if err != nil {
			log.Fatalf("unable to create kubernetes clientset\n%v", err)
		}

		clientset = cs
	}

	return clientset
}

func GetInstalledClusterManifest() []byte {
	api := getClientset().CoreV1().ConfigMaps(config.Config.KubernetesNamespace)
}
