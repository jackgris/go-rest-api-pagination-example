# Go Rest API with pagination example

In this simple example, I will show you how to write a Rest API with pagination and create a Kubernetes cluster using our Docker images.

If you want to run everything with Kubernetes follow the instructions [here](DockerKubernete.md)

Here we have four folders:

- `database`: you will find the Dockerfile configuration for the MySQL database container.
- `back-end`: contain the Go code for the API and the Docker container configuration.
- `fron-end`: is the NextJS code for the front-end that will interact with the API.
- `k8s`: have all the yaml configuration files for running the pods with the three containers.

Also in the `Makefile` you have all the necessary commands to run the Kubernete cluster. And you can build with those commands the Docker images, in that way, you can run it manually running the container or play with the code and the containers.
