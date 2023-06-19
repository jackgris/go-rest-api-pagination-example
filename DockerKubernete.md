# What we need?

First we need [Kubernetes](https://kubernetes.io/releases/download/) and have this tools:

- [Docker](https://docs.docker.com/desktop/)
- [Kubectl, Minikube and Kind](https://kubernetes.io/docs/tasks/tools/)

## First step

Now in the folder fron-end where is the Dockerfile run this command:

```
docker build . -t nextjsapptest
```

If everything is all right you can run this:
```
docker run -p 3000:3000 nextjsapptest
```

And in your browser open the URL: `http://localhost:3000/` and you will see the front end running. After that you can stop the container with:
```
docker container stop "container-name"
```

Note: if you want to know the name of the container running you can use `docker container list`.

## Start our cluster

1- Start Minikube: `minikube start`

2- Start your deplyment: `kubectl apply -f fron-end/deploy.yaml`

3- Verify is running: `kubectl get deployments`

4- Expose the port: `kubectl expose deployment nextjs-app --port 3000`

5- Verify is running: `kubectl get services nextjs-app`

6- Open the port: `minikube service nextjs-app --url`

7- Now you should be able to see again the webpage in the URL: `http://localhost:3000/`


### Fixing problems

#### External IP:
If you have problem getting the local IP and Port you need run this command:
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


Run bash terminal in our MySql database server:

```
kubectl --namespace todos-system exec -it database-759bd947f9-c5rl5 -- bash
```

Port forwarding for our database server:

```
kubectl port-forward service/database 3306:3306 --namespace=todos-system
```
