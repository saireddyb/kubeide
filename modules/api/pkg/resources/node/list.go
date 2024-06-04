package node

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// Add this import
)

// NodeList contains a list of nodes in the cluster.
// type NodeList struct {
// 	Nodes []string `json:"nodes"`
// }

type NodeList struct {
	Nodes []Node `json:"nodes"`
}

type ObjectMetaWithoutManagedFields struct {
	metav1.ObjectMeta `json:",inline"`
	ManagedFields     []metav1.ManagedFieldsEntry `json:"-"`
}

type Node struct {
	ObjectMeta         ObjectMetaWithoutManagedFields `json:"objectMeta"`
	TypeMeta           metav1.TypeMeta                `json:"typeMeta"`
	Name               string                         `json:"name"`
	Labels             map[string]string              `json:"labels"`
	Ready              v1.ConditionStatus             `json:"ready"`
	AllocatedResources NodeAllocatedResources         `json:"allocatedResources"`
}

type NodeAllocatedResources struct {
	CPU              resource.Quantity `json:"cpu"`
	Memory           resource.Quantity `json:"memory"`
	Pods             resource.Quantity `json:"pods"`
	EphemeralStorage resource.Quantity `json:"ephemeralStorage"`
}

func GetNodeList(client kubernetes.Interface) (NodeList, error) {
	fmt.Println("Get Kubernetes Nodes")
	nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting nodes: %v", err)
		return NodeList{}, err
	}
	// return nodes, nil
	return extractNodeNames(nodes), nil

	// return toNodeList(client, nodes.Items), nil
}

func extractNodeNames(nodes *v1.NodeList) NodeList {
	var nodeInfos []Node
	for _, node := range nodes.Items {
		nodeInfo := Node{
			ObjectMeta: ObjectMetaWithoutManagedFields{
				ObjectMeta: metav1.ObjectMeta{
					Name:                       node.ObjectMeta.Name,
					Namespace:                  node.ObjectMeta.Namespace,
					UID:                        node.ObjectMeta.UID,
					ResourceVersion:            node.ObjectMeta.ResourceVersion,
					Generation:                 node.ObjectMeta.Generation,
					CreationTimestamp:          node.ObjectMeta.CreationTimestamp,
					DeletionTimestamp:          node.ObjectMeta.DeletionTimestamp,
					DeletionGracePeriodSeconds: node.ObjectMeta.DeletionGracePeriodSeconds,
					Labels:                     node.ObjectMeta.Labels,
					Annotations:                node.ObjectMeta.Annotations,
					OwnerReferences:            node.ObjectMeta.OwnerReferences,
					Finalizers:                 node.ObjectMeta.Finalizers,
				},
			},
			TypeMeta:           node.TypeMeta,
			Name:               node.Name,
			Labels:             node.Labels,
			Ready:              node.Status.Conditions[0].Status, // assuming the first condition is the "Ready" condition
			AllocatedResources: NodeAllocatedResources{},         // replace with actual allocated resources
		}
		nodeInfos = append(nodeInfos, nodeInfo)
	}
	return NodeList{Nodes: nodeInfos}
}
