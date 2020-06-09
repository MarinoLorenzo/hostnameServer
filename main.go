/*
	Symple application to discover how to create an image container and then deploy the app
	on Kubernetes
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from", r.RemoteAddr) //different from r.Host
	//ok := strconv.Itoa(http.StatusOK)
	fmt.Fprintf(w, "200 OK\n")
	host, _ := os.Hostname()
	fmt.Fprintf(w, "You've hit "+host+"\n")

}

func main() {
	fmt.Println("---Who is the Host APP---")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
