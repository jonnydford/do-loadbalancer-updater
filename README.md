# DigitalOcean Load Balancer Updater
_Update Load Balancer based on Digital Ocean Droplets or Tags_

Designed for use with blue/green deployments, but not limited to this. 

* Linux 32 and 64 bit

## Build From Source
Alternatively, if you have a Go environment configured, you can install the development version from the command line like so:

```
$ go get github.com/jonnydford/do-loadbalancer-updater
```

## Using DigitalOcean Load Balancer Updater

```
/path/to/executable --token DIGITALOCEAN_API_TOKEN --loadbalancer-id LOADBALANCER_ID
```

Your API token must have read and write privileges.

## Variables
``--token``
``--loadbalancer-id`` or ``--loadbalancer-name``
``--droplet-tag`` or ``--dropletIDs``

All of the variables can also be set in your environment path, useful especially for the API token.
