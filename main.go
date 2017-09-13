package dolbupdater

import (
	"flag"
	"github.com/digitalocean/godo"
	"log"
	"os"
)

func main() {
	var apiToken, loadbalancerName, loadbalancerID, dropletTag, dropletIDs string
	var err error

	/// We get all of these from your environment variables if you so choose to set them
	os.Getenv("apiToken")
	os.Getenv("loadbalancerName")
	os.Getenv("loadbalancerID")
	os.Getenv("dropletTag")
	os.Getenv("dropletIDs")

	flag.StringVar(&apiToken, "token", "", "The digitalocean api token, can also export apiToken")
	flag.StringVar(&loadbalancerName, "loadbalancer-name", "", "The name of the loadbalancer")
	flag.StringVar(&loadbalancerID, "loadbalancer-id", "", "The id of the loadbalancer")
	flag.StringVar(&dropletTag, "droplet-tag", "", "The droplet tag name for the loadbalancer")
	flag.StringVar(&dropletIDs, "dropletIDs", "", "The droplet ID for the load balancer")

	flag.Parse()

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
		DropletIDs:          LoadBalancer.DropletIDs,
		Tag:                 LoadBalancer.Tag,
		RedirectHttpToHttps: LoadBalancer.RedirectHttpToHttps,
	}

	updateLoadBalancer(client, loadbalancerID, lb, err)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Load Balancer updated successfully.")
}
