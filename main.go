package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/heptio/clerk/cluster"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get the cluster version
	// version := cluster.Version(clientset)

	// Init our cluster inventory with the version
	// inv := inventory.Cluster{
	// 	Version: version,
	// }
	// log.Println(inv.Version)

	// watch namespaces
	go func() {
		cluster.Namespaces(clientset)
	}()

	// // init nsDeployments var for use in loop
	// nsDeployments := make(map[string][]inventory.Deployment, len(inv.Namespaces))
	// // fetch deployment information
	// // nsDeployments := make(map[string][]inventory.Deployment)
	// for _, ns := range inv.Namespaces {
	// 	nsDeployments[ns] = cluster.Deployments(clientset, ns)
	// }

	// // take our adhoc nsDeployments and add it to the inventory struct
	// inv.Deployments = nsDeployments

	// testing
	// fmt.Printf("%+v", inv.Deployments)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting.")
}
