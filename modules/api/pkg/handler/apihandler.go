package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/saireddyb/kubeide/pkg/client"
	"github.com/saireddyb/kubeide/pkg/resources/node"
	"k8s.io/client-go/kubernetes"
)

type apiHandler struct {
	ClientSet *kubernetes.Clientset
}

func InitializeClient() *kubernetes.Clientset {
	// Sample kubeconfig and context
	home, _ := os.UserHomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	// Get the current context
	context, err := client.GetCurrentContext(kubeconfig)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current context:", context)

	clientset, err := client.Create(kubeconfig, context)
	return clientset
}

// CreateHTTPAPIHandler creates a new HTTP handler that handles all requests to the API of the backend.
func CreateHTTPAPIHandler() (http.Handler, error) {
	mux := http.NewServeMux()
	basePath := "/api/v1"
	clientset := InitializeClient()
	apiHandler := &apiHandler{ClientSet: clientset}

	mux.HandleFunc(basePath+"/nodes", apiHandler.GetNodeList)
	mux.HandleFunc(basePath+"/hello", apiHandler.GetHellos)

	return mux, nil
}

func (apiHandle *apiHandler) GetNodeList(w http.ResponseWriter, r *http.Request) {
	// Get the list of nodes
	nodes, _ := node.GetNodeList(apiHandle.ClientSet)
	// write a sameple hello world to the response
	// fmt.Println(apiHandle.ClientSet)
	// w.Write([]byte("Hello World"))

	// // Write the nodes to the response
	// apiHandler.writeJSON(w, http.StatusOK, nodes)
	json.NewEncoder(w).Encode(nodes)
}

// For /hello path
func (apiHandle *apiHandler) GetHellos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
