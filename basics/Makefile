PROJECT = com.iximiuz.projects.envoy-playground-basics
EXPOSE_SERVICE_A = 8080
EXPOSE_PROMETHEUS = 9090
SERVICE_B_ERROR_RATE=20

pod_infra_container_id = $(shell podman pod inspect -f "{{.InfraContainerID}}" $(1))
pod_ip_address = $(shell podman inspect -f "{{.NetworkSettings.IPAddress}}" $(call pod_infra_container_id,$(1)))

tgreen := $(shell tput setaf 2)
treset := $(shell tput sgr0)


.PHONY: run
run: service-b-pod-create service-b-build service-b-run \
	service-a-pod-create service-a-build service-a-run \
	prometheus-build prometheus-run


.PHONY: test
test:
	for _ in {1..1000}; do curl --silent localhost:${EXPOSE_SERVICE_A}; done | sort | uniq -w 8 -c


.PHONY: service-a-pod-create
service-a-pod-create:
	@echo "${tgreen}>> creating Service A pod${treset}"
	@podman pod create --name service-a-pod \
	   	--label ${PROJECT} \
		--add-host b.service:$(call pod_ip_address,service-b-pod) \
		--publish ${EXPOSE_SERVICE_A}:80


.PHONY: service-a-build
service-a-build: service-a/Dockerfile
	@echo "${tgreen}>> building Service A image${treset}"
	@podman build -t ${PROJECT}/service-a -f service-a/Dockerfile


.PHONY: service-a-run
service-a-run:
	@echo "${tgreen}>> starting Service A${treset}"
	@podman run -d --rm --pod service-a-pod \
		--label ${PROJECT} ${PROJECT}/service-a


.PHONY: service-b-pod-create
service-b-pod-create:
	@echo "${tgreen}>> creating Service B pod${treset}"
	@podman pod create --name service-b-pod --label ${PROJECT}


.PHONY: service-b-build
service-b-build: service-b/Dockerfile
	@echo "${tgreen}>> building Service B image${treset}"
	@podman build -t ${PROJECT}/service-b -f service-b/Dockerfile


.PHONY: service-b-run
service-b-run:
	@echo "${tgreen}>> starting Service B${treset}"
	@podman run -d --rm --pod service-b-pod \
		-e ERROR_RATE=${SERVICE_B_ERROR_RATE} \
		--label ${PROJECT} ${PROJECT}/service-b


.PHONY: prometheus-build
prometheus-build: prometheus/Dockerfile
	@echo "${tgreen}>> building prometheus image${treset}"
	@podman build -t ${PROJECT}/prometheus -f prometheus/Dockerfile


.PHONY: prometheus-run
prometheus-run:
	@echo "${tgreen}>> starting prometheus service${treset}"
	@podman run -d \
	   	--label ${PROJECT} \
		--publish ${EXPOSE_PROMETHEUS}:9090 \
		--add-host a.service:$(call pod_ip_address,service-a-pod) \
		--add-host b.service:$(call pod_ip_address,service-b-pod) \
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

