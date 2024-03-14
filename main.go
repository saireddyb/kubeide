package main

import (
	"context"
	"fmt"
	"os"

	"github.com/saireddyb/kubeide/pkg/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	// Sample kubeconfig and context
	kubeconfig := "/home/sainath_reddy/.kube/config"
	
	// Get the current context
	context, err := client.GetCurrentContext(kubeconfig)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current context:", context)

	


	clientset, err := client.Create(kubeconfig, context)
	namespaces, err := ListNamespaces(clientset)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	for _, namespace := range namespaces.Items {
		fmt.Println(namespace.Name)
	}
	fmt.Printf("Total namespaces: %d\n", len(namespaces.Items))

}


func ListNamespaces(client kubernetes.Interface) (*v1.NamespaceList, error) {
	fmt.Println("Get Kubernetes Namespaces")
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting namespaces: %v", err)
		return nil, err
	}
	return namespaces, nil
}

