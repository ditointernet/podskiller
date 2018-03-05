package main

import (
	"flag"
	"fmt"
	"regexp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var regex = flag.String("regex", "", "Regex to filter pods deletion")
	var namespace = flag.String("namespace", "default", "Kubernetes namespace")
	flag.Parse()

	if len(*regex) == 0 {
		panic("regex is required")
	}
	re := regexp.MustCompile(*regex)

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(*namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		name := pod.GetName()
		if re.MatchString(name) {
			fmt.Printf("Deleting pod %s...", name)
			err := clientset.CoreV1().Pods(*namespace).Delete(name, &metav1.DeleteOptions{})
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Done!")
		}
	}
}
