# DigitalOcean Load Balancer Updater
_Update Load Balancer based on Digital Ocean Droplets or Tags_

Designed for use with blue/green deployments, but not limited to this. 

## Build From Source

```
$ go get github.com/jonnydford/do-loadbalancer-updater
```

## Releases
In the future I'll be releasing a precompiled binary for you to download. 

## Using DigitalOcean Load Balancer Updater

```
/path/to/executable --token DIGITALOCEAN_API_TOKEN --loadbalancer-id LOADBALANCER_ID --droplet-tag TAGNAME --region AMS2
```

Your API token must have read and write privileges.

## Variables
``--token``
``--loadbalancer-id`` or ``--loadbalancer-name``
``--droplet-tag``

All of the variables can also be set in your environment path, useful especially for the API token.
