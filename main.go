package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/digitalocean/godo"
)

func main() {
	var apiToken, loadbalancerName, loadbalancerID, dropletTag, region string
	var err error

	/// We get all of these from your environment variables if you so choose to set them
	os.Getenv("apiToken")
	os.Getenv("loadbalancerName")
	os.Getenv("loadbalancerID")
	os.Getenv("dropletTag")
	os.Getenv("dropletIDs")
	os.Getenv("region")

	flag.StringVar(&apiToken, "token", "", "The digitalocean api token, can also export apiToken")
	flag.StringVar(&loadbalancerName, "loadbalancer-name", "", "The name of the loadbalancer")
	flag.StringVar(&loadbalancerID, "loadbalancer-id", "", "The id of the loadbalancer")
	flag.StringVar(&dropletTag, "droplet-tag", "", "The droplet tag name for the loadbalancer")
	flag.StringVar(&region, "region", "", "Region name, such as ams2")

	flag.Parse()

	if apiToken == "" {
		log.Fatal("You must specify the --token")
	}

	if loadbalancerName == "" && loadbalancerID == "" {
		log.Fatal("You must specify the --loadbalancer-name or the --loadbalancer-id")
	}

	if dropletTag == "" {
		log.Fatal("You must specify the --droplet-tag")
	}

	client := newClient(apiToken)

	var LoadBalancer *godo.LoadBalancer
	if loadbalancerID != "" {
		LoadBalancer, err = findLoadBalancerByID(client, loadbalancerID)
	} else {
		LoadBalancer, err = findLoadBalancerByName(client, loadbalancerName)
	}

	if err != nil {
		log.Fatal(err)
	}

	lb := &godo.LoadBalancerRequest{
		Name:                LoadBalancer.Name,
		Algorithm:           LoadBalancer.Algorithm,
		ForwardingRules:     LoadBalancer.ForwardingRules,
		HealthCheck:         LoadBalancer.HealthCheck,
		StickySessions:      LoadBalancer.StickySessions,
		Tag:                 LoadBalancer.Tag,
		RedirectHttpToHttps: LoadBalancer.RedirectHttpToHttps,
	}

	lb.Region = strings.ToLower(region)
	lb.Tag = dropletTag

	updateLoadBalancer(client, LoadBalancer.ID, lb, err)

	log.Println("Load Balancer updated successfully.")
}
