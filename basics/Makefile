PROJECT = com.iximiuz.projects.envoy-playground-basics
EXPOSE_FRONTEND = 8080
EXPOSE_METRICS = 9090
BACKEND_ERROR_RATE=30

pod_infra_container_id = $(shell podman pod inspect -f "{{.InfraContainerID}}" $(1))
pod_ip_address = $(shell podman inspect -f "{{.NetworkSettings.IPAddress}}" $(call pod_infra_container_id,$(1)))

tgreen := $(shell tput setaf 2)
treset := $(shell tput sgr0)


.PHONY: run
run: backend-pod-create backend-svc-build backend-svc-run \
	frontend-pod-create frontend-svc-build frontend-svc-run \
	prometheus-build prometheus-run


.PHONY: backend-pod-create
backend-pod-create:
	@echo "${tgreen}>> creating backend pod${treset}"
	@podman pod create --name backend-pod --label ${PROJECT}


.PHONY: backend-svc-build
backend-svc-build: backend-svc/Dockerfile
	@echo "${tgreen}>> building backend service image${treset}"
	@podman build -t ${PROJECT}/backend-svc -f backend-svc/Dockerfile


.PHONY: backend-svc-run
backend-svc-run:
	@echo "${tgreen}>> starting backend service${treset}"
	@podman run -d --rm --pod backend-pod \
		-e ERROR_RATE=${BACKEND_ERROR_RATE} \
		--label ${PROJECT} ${PROJECT}/backend-svc


.PHONY: frontend-pod-create
frontend-pod-create:
	@echo "${tgreen}>> creating frontend pod${treset}"
	@podman pod create --name frontend-pod \
	   	--label ${PROJECT} \
		--add-host backend.service:$(call pod_ip_address,backend-pod) \
		--publish ${EXPOSE_FRONTEND}:80


.PHONY: frontend-svc-build
frontend-svc-build: frontend-svc/Dockerfile
	@echo "${tgreen}>> building frontend service image${treset}"
	@podman build -t ${PROJECT}/frontend-svc -f frontend-svc/Dockerfile


.PHONY: frontend-svc-run
frontend-svc-run:
	@echo "${tgreen}>> starting frontend service${treset}"
	@podman run -d --rm --pod frontend-pod \
		--label ${PROJECT} ${PROJECT}/frontend-svc


.PHONY: prometheus-build
prometheus-build: prometheus/Dockerfile
	@echo "${tgreen}>> building prometheus image${treset}"
	@podman build -t ${PROJECT}/prometheus -f prometheus/Dockerfile


.PHONY: prometheus-run
prometheus-run:
	@echo "${tgreen}>> starting prometheus service${treset}"
	@podman run -d \
	   	--label ${PROJECT} \
		--publish ${EXPOSE_METRICS}:9090 \
		--add-host frontend.service:$(call pod_ip_address,frontend-pod) \
		--add-host backend.service:$(call pod_ip_address,backend-pod) \
		${PROJECT}/prometheus


.PHONY: kill-containers
kill-containers:
	@echo "${tgreen}>> killing containers${treset}"
	@podman ps -q -f label=${PROJECT} | xargs -I % podman kill %


.PHONY: remove-containers
remove-containers:
	@echo "${tgreen}>> removing containers${treset}"
	@podman ps -qa -f label=${PROJECT} | xargs -I % podman rm %


.PHONY: remove-pods
remove-pods:
	@echo "${tgreen}>> removing pods${treset}"
	@podman pod ls -q -f label=${PROJECT} | xargs -I % podman pod rm %


.PHONY: remove-images
remove-images:
	@echo "${tgreen}>> removing images${treset}"
	@podman images -q ${PROJECT}/* | xargs -I % podman rmi %


.PHONY: clean
clean: kill-containers remove-pods remove-containers remove-images

