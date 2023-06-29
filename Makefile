# ==============================================================================
# Define dependencies

KIND            := kindest/node:v1.27.3
KIND_CLUSTER    := goscrapy-starter-cluster
NAMESPACE       := todos-system
NAMESPACE_DB		:= db-ns
NAMESPACE_API		:= todos-ns
APP             := todosapi
VERSION         := 0.0.1
NAME_DOCKER_HUB := jackgris
BASE_IMAGE_NAME := pagination
SERVICE_NAME    := todos-api
SERVICE_IMAGE   := $(NAME_DOCKER_HUB)/$(BASE_IMAGE_NAME)-$(SERVICE_NAME):$(VERSION)
DATABASE_NAME   := todos-mysql
DATABASE_IMAGE  := $(NAME_DOCKER_HUB)/$(BASE_IMAGE_NAME)-$(DATABASE_NAME):$(VERSION)
FRONT_END_NAME  := todos-front-end
FRONT_END_IMAGE := $(NAME_DOCKER_HUB)/$(BASE_IMAGE_NAME)-$(FRONT_END_NAME):$(VERSION)

# ==============================================================================
# Running from within k8s/kind

dev-up-local:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

	kubectl apply -f namespace.yaml

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

# ==============================================================================
# Building containers

all: database-mysql api-todo front-end-app

service:
	cd back-end; \
	docker build \
		-f ./Dockerfile \
		-t $(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

database-mysql:
	docker build \
		-f database/Dockerfile \
		-t $(DATABASE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	  .

front-end-app:
	docker build \
		--build-arg CONFIG_API_HOST='todosapi.todos-ns.svc.cluster.local:3000' \
    --build-arg CONFIG_SERVER_PORT='8080' \
		-f fron-end/Dockerfile \
		-t $(FRONT_END_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	  .

# ------------------------------------------------------------------------------
# Load and manage images
dev-load-database:
	kubectl create -f k8s/database/mysql-ns.yaml
	kubectl create -f k8s/database/configmap-mysql.yaml
	kubectl create -f k8s/database/secret-mysql.yaml
	kubectl create -f k8s/database/mysql-volume-pvc.yaml
	kubectl create -f k8s/database/deployment-mysql.yaml
	kubectl create -f k8s/database/service-mysql.yaml

dev-load-api:
	kubectl create -f k8s/api/todos-ns.yaml
	kubectl create -f k8s/api/secret-todos.yaml
	kubectl create -f k8s/api/configmap-todos.yaml
	kubectl create -f k8s/api/deployment-todos.yaml
	kubectl create -f k8s/api/service-todos.yaml

dev-load-front:
	kubectl create -f k8s/front/front-end-ns.yaml
	kubectl create -f k8s/front/configmap-front-end.yaml
	kubectl create -f k8s/front/deployment-front-end.yaml
	kubectl create -f k8s/front/service-front-end.yaml

dev-load: dev-load-database dev-load-api

dev-port-forward:
	kubectl port-forward -n todos-ns svc/todosapi 3000 --namespace=todos-ns

dev-status:
	kubectl get nodes -o wide --namespace=$(NAMESPACE_DB)
	kubectl get nodes -o wide --namespace=$(NAMESPACE_API)
	kubectl get svc -o wide --namespace=$(NAMESPACE_DB)
	kubectl get svc -o wide --namespace=$(NAMESPACE_API)
	kubectl get pods -o wide --watch --all-namespaces

dev-describe:
	kubectl describe nodes --namespace=$(NAMESPACE_DB)
	kubectl describe nodes --namespace=$(NAMESPACE_API)
	kubectl describe svc --namespace=$(NAMESPACE_DB)
	kubectl describe svc --namespace=$(NAMESPACE_API)

dev-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE_API) $(APP)

dev-describe-todos:
	kubectl describe pod --namespace=$(NAMESPACE_API) -l app=$(APP)

dev-pods:
	kubectl get pods --all-namespaces

# ------------------------------------------------------------------------------
