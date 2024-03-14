package client

import "k8s.io/client-go/tools/clientcmd"

func getCurrentContext(kubeconfig string) (string, error) {
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		return "", err
	}
	currentContext := config.CurrentContext
	return currentContext, nil
}
