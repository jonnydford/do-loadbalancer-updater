package dolbupdater

import (
	"context"
	"errors"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
	"strings"
)

type tokenSource struct {
	AccessToken string
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func newClient(token string) *godo.Client {
	tokenSource := &tokenSource{
		AccessToken: token,
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	return client
}

func findLoadBalancerByName(client *godo.Client, name string) (loadbalancer *godo.LoadBalancer, err error) {
	ctx := context.TODO()
	options := &godo.ListOptions{
		PerPage: 200,
	}

	loadbalancers, _, err := client.LoadBalancers.List(ctx, options)
	if err != nil {
		return
	}

	for _, loadbalancer := range loadbalancers {
		if !strings.EqualFold(loadbalancer.Name, name) {
			continue
		}

		return &loadbalancer, nil
	}

	err = errors.New("loadbalancer Not Found")
	return
}

func findLoadBalancerByID(client *godo.Client, lbID string) (loadbalancer *godo.LoadBalancer, err error) {
	ctx := context.TODO()

	loadbalancer, _, err = client.LoadBalancers.Get(ctx, lbID)
	return
}

func updateLoadBalancer(client *godo.Client, lbID string, lb *godo.LoadBalancerRequest, err error) {
	ctx := context.TODO()

	_, _, err = client.LoadBalancers.Update(ctx, lbID, lb)
	return
}
