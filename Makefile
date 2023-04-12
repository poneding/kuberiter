IMG_TAG ?= latest

.PHONY: all
all: docker-build

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: manifests
manifests: ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	controller-gen rbac:roleName=kuberiter-sa crd webhook paths="./..." output:crd:artifacts:config=config/crds
	find config/crds -type f -exec cat {} \; > deploy/crds.yaml
	cp -r config/crds/ chart/crds

.PHONY: generate
generate: ## Generate code containing DeepCopy Functions, Clientset, Informers and Listers.
	./hack/update-codegen.sh

.PHONY: install
install: manifests ## Install crds into the K8s cluster specified in ~/.kube/config.
	kubectl apply -f config/crds

.PHONY: uninstall
uninstall: ## Uninstall crds from the K8s cluster specified in ~/.kube/config.
	kubectl delete -f config/crds

.PHONY: deploy
deploy: manifests ## Deploy kuberiter to the K8s cluster specified in ~/.kube/config.
	kubectl apply -f deploy/crds/
	kubectl apply -f deploy/manifests.yaml

.PHONY: docker-build
docker-build: docker-build-api-server docker-build-controller-manager ## Build docker images

.PHONY: docker-build-api-server
docker-build-api-server: ## Build kuberiter-api-server docker image
	docker build -t poneding/kuberiter-api-server:$(IMG_TAG) -f ./build/api-server/Dockerfile . 

.PHONY: docker-build-controller-manager
docker-build-controller-manager: ## Build kuberiter-controller-manager docker image
	docker build -t poneding/kuberiter-controller-manager:$(IMG_TAG) -f ./build/controller-manager/Dockerfile .

.PHONY: docker-push
docker-push: docker-push-api-server docker-push-controller-manager ## Push docker images

.PHONY: docker-push-api-server
docker-push-api-server: ## Push kuberiter-api-server docker image
	docker push poneding/kuberiter-api-server:$(IMG_TAG)

.PHONY: docker-push-controller-manager
docker-push-controller-manager: ## Push kuberiter-controller-manager docker image
	docker push poneding/kuberiter-controller-manager:$(IMG_TAG)
