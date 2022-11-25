package main

import (
	"context"
	"flag"
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeConfig := flag.String("kubeconfig", "/root/config", "")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
		}
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
	}
	pods, err := clientSet.CoreV1().Pods("default").List(context.Background(), metaV1.ListOptions{})
	if err != nil {
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.String())
		fmt.Println(pod.GetName())
	}

	 CreateController(clientSet)

}
