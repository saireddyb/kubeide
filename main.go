package main

import (
	"fmt"

	"github.com/saireddyb/kubeide/pkg/client"
)

func main() {
	// Sample kubeconfig and context
	kubeconfig := "/home/sainath_reddy/.kube/config"
	
	// Get the current context
	context, _ := client.getCurrentContext(kubeconfig)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	fmt.Println("Current context:", context)

}



