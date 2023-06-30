# What do we need?

First we need [Kubernetes](https://kubernetes.io/releases/download/) and this tools installed:

- [Docker](https://docs.docker.com/desktop/)
- [Kubectl, Minikube and Kind](https://kubernetes.io/docs/tasks/tools/)

## Running everything in Kubernetes

If you want to see everything running, I created a cluster with Kind, you can use these commands on the root folder of the project to run it. (You need `make` for that, if you don't want to install make, you can run the commands you can see in the Makefile)

#### Start cluster

```
make dev-up-local
```

#### Start containers

```
 make dev-load
```

#### Wait for all pods

With this command:
```
 make dev-pods
```

You can see if everything is running, that make a few minutes because the API pod needs to wait at the MySql pod before starting.

### Port Forwarding

In two different terminals, you need to run:

```
make dev-port-forward-api
```
and
```
make dev-port-forward-front
```

### See the app running

Now you only need to go to this URL: `http://localhost:8080`

If you running that in Firefox on Linux, you could have troubles with CORS, to fix that you need to change the configuration of Firefox or use another browser like `Brave`

### Stop everything

With this command, you will stop all the pods, because you will stop the cluster.
```
make dev-down
```

## Running everything using only Docker

### Build images

This command will create the three images that you need to use: (front end, api back end, and database)
```
make all
```
### After that you can run the container

You need to follow this pattern, If everything is all right you can run this:
```
docker run --name newname -p 8080:8080 myimage
```

Where `3000` is the port described in the Dockerfile, and myimage is the name of the images created with the command `make all`. (`newname` is the name of the container).

After doing that the first time, you can run the container with these commands:

```
docker container start newname
```

And in your browser open the URL: `http://localhost:8080/v1/todos` and you will see the front end running. After that, you can stop the container with:
```
docker container stop "container-name"
```

Note: if you want to know the name of the container running you can use `docker container list`.

### Fixing problems

#### External IP:
If you have a problem getting the local IP and Port you need to run this command:
```
kubectl port-forward service/<service-name> <local-port>:<service-port>
```
#### ErrImageNeverPull:
When trying to use a local container image maybe you receive this error, for fix that run this in your terminal `minikube docker-env` and follow the instructions, something like this at the end:

```
# To point your shell to minikube’s docker-daemon, run:
# eval $(minikube -p minikube docker-env)
```

So, run that command: `eval $(minikube -p minikube docker-env)` and build your container again.


### Accessing our database

Run bash terminal in our MySql database server:

```
kubectl exec -it -n db-ns mysql-65577f9f69-cqpwk  -- bash
```

Where `mysql-65577f9f69-cqpwk` is the name of the Pod that runs our database. After that, you can use MySQL command line with:

```
mysql -u root -p
```
Remember the password is in the secret configuration `yaml` file or in the Docker environment variables. (in this case, the super safe `1234` password)


### Accessing our API Todo service

Port forwarding for our database server:

```
kubectl port-forward -n todos-ns svc/todosapi 3000 --namespace=todos-ns
```

You also have these commands in the Makefile.
