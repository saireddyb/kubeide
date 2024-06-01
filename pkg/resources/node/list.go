package node

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeList(client kubernetes.Interface) (*v1.NodeList, error) {
	fmt.Println("Get Kubernetes Nodes")
	nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting nodes: %v", err)
		return nil, err
	}
	return nodes, nil
}
