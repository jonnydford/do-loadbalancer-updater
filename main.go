package main

import (
	"github.com/digitalocean/godo"
	"log"
    "flag"
    "os"
)

func main() {
	var apiToken, loadbalancerName, loadbalancerID, dropletTag, dropletIDs string
    var err error

	flag.StringVar(&apiToken, "token", "", "The digitalocean api token, can also export apiToken")
	flag.StringVar(&loadbalancerName, "loadbalancer-name", "", "The name of the loadbalancer")
	flag.StringVar(&loadbalancerID, "loadbalancer-id", "", "The id of the loadbalancer")
	flag.StringVar(&dropletTag, "droplet-tag", "", "The droplet tag name for the loadbalancer")
	flag.StringVar(&dropletIDs, "dropletIDs", "", "The droplet ID for the load balancer")

	flag.Parse()

/// We get all of these from your environment variables if you so choose to set them
    os.Getenv("apiToken")
    os.Getenv("loadbalancerName")
    os.Getenv("loadbalancerID")
    os.Getenv("dropletTag")
    os.Getenv("dropletIDs")

	if apiToken == "" {
		log.Fatal("You must specify the --token")
	}

	if loadbalancerName == "" && loadbalancerID == "" {
		log.Fatal("You must specify the --loadbalancer-name or the --loadbalancer-id")
	}

	if dropletTag == "" && dropletIDs == "" {
		log.Fatal("You must specify the --droplet-tag or the --dropletIDs")
	}

	if dropletTag != "" && dropletIDs != "" {
		log.Fatal("Just specify one of --droplet-tag or --dropletIDs")
	}

	client := newClient(apiToken)

	var loadbalancer *godo.LoadBalancer
	if loadbalancerID != "" {
		loadbalancer, err = findLoadBalancerByID(client, loadbalancerID)
	} else {
		loadbalancer, err = findLoadBalancerByName(client, loadbalancerName)
	}

	if err != nil {
		log.Fatal(err)
	}

	lb := &godo.LoadBalancerRequest{
		Name:                LoadBalancer.Name,
		Algorithm:           LoadBalancer.Algorithm,
		Region:              LoadBalancer.Region,
		ForwardingRules:     LoadBalancer.ForwardingRules,
		HealthCheck:         LoadBalancer.HealthCheck,
		StickySessions:      LoadBalancer.StickySessions,
		DropletIDs:          LoadBalancer.DropletIDs,
		Tag:                 LoadBalancer.Tag,
		RedirectHttpToHttps: LoadBalancer.RedirectHttpToHttps,
	}

	err = updateLoadBalancer(client, loadbalancerID, lb)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Load Balancer updated successfully.")
}
