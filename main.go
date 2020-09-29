package main

import (
	"context"
	"flag"
	"fmt"
	"regexp"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var regex = flag.String("regex", "", "Regex to filter pods deletion")
	var namespace = flag.String("namespace", "default", "Kubernetes namespace")
	var delay = flag.Duration("delay", 0, "Delay between pod deletions")
	flag.Parse()

	if len(*regex) == 0 {
		panic("regex is required")
	}
	re := regexp.MustCompile(*regex)

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Starting pods killer to kill all \"%s\" pods with delay of %s\n", *regex, delay.String())

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()

	pods, err := clientset.CoreV1().Pods(*namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		name := pod.GetName()
		if re.MatchString(name) {
			if delay != nil {
				time.Sleep(*delay)
			}
			fmt.Printf("Deleting pod %s...\n", name)
			err := clientset.CoreV1().Pods(*namespace).Delete(ctx, name, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Done!")
		}
	}
}
