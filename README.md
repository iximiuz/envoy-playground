# Envoy Playground

Container-based playground showing various capabilities of service proxy pattern.

## Requirements

All the scenarios have been tested on `CentOS 8` (see the Vagrantfile at the root of the repository) using `podman 2.0.4`.

## Basics

The environment creates two services. The frontend service _A_ call the faulty upstream service _B_. Envoy is used as a sidecar to mitigate the erroneous behavior of the upstream service.

Run it:

```bash
# Build images
$ sudo make build

# Run A <-> B scenario
$ sudo make run-direct

# Pour some traffic, notice error ratio
$ sudo make traffic
> for _ in {1..1000}; do curl --silent localhost:8080; done | sort | uniq -w 24 -c
>    1000
>     196 Service A: upstream failed with: HTTP 500 - Service B: Ooops... nounce 1013789005
>     804 Service A: upstream responded with: Service B: Yay! nounce 1003014939

# Stop containers
$ sudo make stop

# Run A <-> envoy <-> B scenario
$ sudo make run-envoy

# Pour some traffic, compare error ratio
$ sudo make traffic
> for _ in {1..1000}; do curl --silent localhost:8080; done | sort | uniq -w 24 -c
>    1000
>       9 Service A: upstream failed with: HTTP 500 - Service B: Ooops... nounce 1263663296
>     991 Service A: upstream responded with: Service B: Yay! nounce 1003014939

# Clean up
$ sudo make clean
```
