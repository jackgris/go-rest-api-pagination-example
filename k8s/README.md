# How are divided the files configuration

Inside this folder, you will find the configuration files for our Kubernete cluster. Every folder contains a maximum of six files for every pods, these files will be called with these prefixes or suffixes.

- `configmap`: A ConfigMap is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.  A ConfigMap allows you to decouple environment-specific configuration from your container images, so that your applications are easily portable. [More info](https://kubernetes.io/docs/concepts/configuration/configmap/)
  
- `deployment`: A Deployment provides declarative updates for Pods and ReplicaSets.  You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments. [More info](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
  
- `ns`: In Kubernetes, namespaces provides a mechanism for isolating groups of resources within a single cluster. Names of resources need to be unique within a namespace, but not across namespaces. Namespace-based scoping is applicable only for namespaced objects (e.g. Deployments, Services, etc) and not for cluster-wide objects (e.g. StorageClass, Nodes, PersistentVolumes, etc). [More info](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
  
- `volume`: On-disk files in a container are ephemeral, which presents some problems for non-trivial applications when running in containers. One problem occurs when a container crashes or is stopped. Container state is not saved so all of the files that were created or modified during the lifetime of the container are lost. During a crash, kubelet restarts the container with a clean state. Another problem occurs when multiple containers are running in a Pod and need to share files. It can be challenging to setup and access a shared filesystem across all of the containers. The Kubernetes volume abstraction solves both of these problems. Familiarity with Pods is suggested. [More info](https://kubernetes.io/docs/concepts/storage/volumes/)
  
- `secret`: A Secret is an object that contains a small amount of sensitive data such as a password, a token, or a key. Such information might otherwise be put in a Pod specification or in a container image. Using a Secret means that you don't need to include confidential data in your application code.
Because Secrets can be created independently of the Pods that use them, there is less risk of the Secret (and its data) being exposed during the workflow of creating, viewing, and editing Pods. Kubernetes, and applications that run in your cluster, can also take additional precautions with Secrets, such as avoiding writing secret data to nonvolatile storage.
Secrets are similar to ConfigMaps but are specifically intended to hold confidential data. [More info](https://kubernetes.io/docs/concepts/configuration/secret/)

- `service`: In Kubernetes, a Service is a method for exposing a network application that is running as one or more Pods in your cluster.
A key aim of Services in Kubernetes is that you don't need to modify your existing application to use an unfamiliar service discovery mechanism. You can run code in Pods, whether this is a code designed for a cloud-native world, or an older app you've containerized. You use a Service to make that set of Pods available on the network so that clients can interact with it. [More info](https://kubernetes.io/docs/concepts/services-networking/service/)

So you can modify these files to change the configuration of the cluster.
