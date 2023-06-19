# ==============================================================================
# Define dependencies

KIND            := kindest/node:v1.26.3
KIND_CLUSTER    := goscrapy-starter-cluster
NAMESPACE       := todos-system
APP             := todos
VERSION         := 0.0.1
BASE_IMAGE_NAME := pagination
SERVICE_NAME    := todos-api
SERVICE_IMAGE   := jackgris/$(BASE_IMAGE_NAME)-$(SERVICE_NAME):$(VERSION)
DATABASE_NAME   := todos-mysql
DATABASE_IMAGE  := jackgris/$(BASE_IMAGE_NAME)-$(DATABASE_NAME):$(VERSION)


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

all: database-mysql api-todo

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

# ------------------------------------------------------------------------------
# Load and manage images

dev-load:
	kubectl apply -f database/mysql-pv.yaml
	kubectl apply -f database/mysql-deployment.yaml

# We need to run our $(SERVICE_IMAGE) in --name $(KIND_CLUSTER)

dev-datatbase-client:
	kubectl run -it --rm --image=mysql:8.0.33 --restart=Never mysql-client -- mysql -h mysql -ppassword

dev-apply:
	kubectl rollout status --namespace=$(NAMESPACE) --watch --timeout=120s sts/database

dev-restart:
	kubectl rollout restart deployment $(APP) --namespace=$(NAMESPACE)

dev-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-describe:
	kubectl describe nodes
	kubectl describe svc

dev-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE) $(APP)

dev-describe-sales:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(APP)

# ------------------------------------------------------------------------------

