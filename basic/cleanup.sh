#!/usr/bin/env bash

echo "Kill running containers"
podman kill $(podman ps -q)

echo "Remove pods"
podman pod rm $(podman pod ps -q)

echo "Remove remaining containers"
podman rm $(podman ps -qa)
