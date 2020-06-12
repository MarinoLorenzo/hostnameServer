/*

	Start a simple http server listening on port 80. Every time is hit the server responds
	with the pod and node name where is running

*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//in-cluster config because the app runs in a pod
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	//create clientset to interact with API
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//the name of the current pod is in the environment variable HOSTNAME
	podName, _ := os.Hostname()
	//get the current pod. This requires default service account view permissions
	currentPod, err := clientset.CoreV1().Pods("default").Get(podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	//get node of the current pod
	podNode := currentPod.Spec.NodeName
	fmt.Println("Received request from", r.RemoteAddr) //different from r.Host
	fmt.Fprintf(w, "200 OK\n")
	fmt.Fprintf(w, "You've hit "+podName+" on node: "+podNode+"\n")

}

func main() {

	fmt.Println("---Which pod? Which node?---")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
