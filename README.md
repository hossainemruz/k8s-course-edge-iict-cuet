# Session 3: Accessing an Application on Kubernetes

In this session, we will learn how we can access an application running inside a Kubernetes cluster.

**Theory:**

- Introduction to Service
- Understanding DNS resolution in Kubernetes
- Introduction to Ingress and Ingress Controller

**Hands On:**

- Accessing application using port-forwarding
- Accessing application using NodePort Service
- Accessing application using LoadBalancer Service
- Pod to Pod communication using ClusterIP Service
- Load balancing using  Service
- Accessing application using Ingress

## Accessing application using port-forwarding

Kubernetes port-forwarding let you create a temporary connection between your local machine and a pod.

**Syntax:**

```bash
kubectl port-forward <pod-name> <local-port>:<pod-port>
```

If you omit the local port, Kubernetes will automatically assign a port for you.

**Example:**

At first, let's deploy our sample web server:

```bash
kubectl apply -f manifests/sample-server.yaml
```

Once the pod is running, run the following command to forward the port:

```bash
# specify a local port. it does not have to be the same as the pod port
kubectl port-forward <pod name> 8085:8080

# or let Kubernetes chose a local port for you
kubectl port-forward <pod name> :8080
```

Now, you can access the web server by using `http://localhost:<local port>` URL.

```bash
curl http://localhost:8085/hello
```

**Resources:**

- [Use Port Forwarding to Access Applications in a Cluster](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/)
- [What Is Kubectl Port-Forward and How Does It Work?](https://kodekloud.com/blog/port-forwarding-kubernetes/)
- [How does kubectl port-forward work?](https://dumlutimuralp.medium.com/how-kubectl-port-forward-works-79d0fbb16de3)
- [Kubectl PortForward Examples - How to PortForward in K8s](https://www.middlewareinventory.com/blog/kubectl-port-forward/)

## Accessing application using NodePort Service

NodePort service let you expose your application on a certain port on each node in the cluster.

Let's create a NodePort service for our sample web server:

```bash
kubectl apply -f manifests/node-port-service.yaml
```

Describe the Service to get details about the NodePort assigned to the service:

```bash
kubectl describe service node-port-service
```

Now, you can access the sample server using `http://<node-ip>:<node-port>` URL.

To get your node IPs, run the following command:

```bash
kubectl get nodes -o wide
```

You will see the IP addresses of the nodes in `INTERNAL-IP` column. You can use any of these node IP address to access the service.

```bash
curl http://<node-ip>:<node-port>/hello
```

## Accessing application using LoadBalancer Service

LoadBalancer service let you expose your application outside the cluster using a cloud provider's load balancer.

At first, install and run the [cloud-provider-kind](https://github.com/kubernetes-sigs/cloud-provider-kind) to simulate a cloud provider in the local Kubernetes cluster:

```bash
# install the cloud-provider-kind
go install sigs.k8s.io/cloud-provider-kind@latest

# run the cloud-provider-kind
./cloud-provider-kind
```

Then, create a LoadBalancer service for our sample web server:

```bash
kubectl apply -f manifests/load-balancer-service.yaml
```

Now, run the following command and wait for the `EXTERNAL-IP` to be assigned to the service:

```bash
watch -n 1 kubectl get service load-balancer-service
```

Once the EXTERNAL-IP is assigned, you can access the sample server using `http://<external-ip>:<port>` URL.

```bash
curl http://<external-ip>:80/hello
```

## Pod to Pod communication using ClusterIP Service

At first, create a ClusterIP service for our sample web server:

```bash
kubectl apply -f manifests/cluster-ip-service.yaml
```

Now, any application running inside cluster can access the server using following URLs:

```bash
# From same namespace
curl http://cluster-ip-service:<port>

# From different namespace. Here, `default` is the namespace where the service is running
curl http://cluster-ip-service.default.svc:<port>
```

Let's create the sample ping client to test the communication. Set the `PING_URL` environment variable to `http://cluster-ip-service:80/hello`.

```bash
kubectl apply -f manifests/sample-ping-client.yaml
```

Now, check the logs of the ping client pod to see if it can access the sample server:

```bash
kubectl logs -l app=ping-client -f
```

## Load balancing using Service

Service in Kubernetes provides load balancing for the pods. It distributes the traffic among the pods that are selected by the service.

Let's scale the sample server to 3 replicas:

```bash
kubectl scale deployment sample-web-server --replicas=3
```

Now, let's use the NodePort service to send periodic request from our local machine to the sample server:

```bash
while true; do
  curl http://<node ip>:<node port>/hello
  sleep 1
done
```

You will see that the requests are distributed among the pods.

## Accessing application using Ingress

Ingress let's you route external traffic to your services inside the cluster based on host, path etc.

At first, let's apply a label to our control-plane node so that ingress controller can run on it:

```bash
kubectl label node kind-control-plane ingress-ready=true
```

Now, install the Nginx Ingress Controller:

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

Wait for the ingress controller to be running:

```bash
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=300s
```

If you list services in the `ingress-nginx` namespace, you will see a NodePort type service named `ingress-nginx-controller`. We will use this service to access the ingress.

```bash
kubectl get service -n ingress-nginx
```

Let's deploy two sample applications `foo` and `bar`:

```bash
kubectl apply -f manifests/app-foo.yaml
kubectl apply -f manifests/app-bar.yaml
```

This will create two deployments with our sample web server image and ClusterIP service for for them. Verify that the services have been created:

```bash
kubectl get service
```

Now, create the Ingress resource:

```bash
kubectl apply -f manifests/ingress.yaml
```

Finally, we can access the sample applications using following URLs.

```bash
# access foo application
curl http://<any node ip>:<node port>/foo/hello

# access bar application
curl http://<any node ip>:<node port>/bar/hello
```

## Resources

- [Use a Service to Access an Application in a Cluster](https://kubernetes.io/docs/tasks/access-application-cluster/service-access-application-cluster/)
- [The difference between ClusterIP, NodePort, and LoadBalancer Kubernetes services](https://octopus.com/blog/difference-clusterip-nodeport-loadbalancer-kubernetes)
- [Kubernetes NodePort vs LoadBalancer vs Ingress? When should I use what?](https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0)
- [3 Ways to Expose Applications Running in Kubernetes Cluster to Public Access](https://medium.com/@seanlinsanity/how-to-expose-applications-running-in-kubernetes-cluster-to-public-access-65c2fa959a3b)
- [Ingress Documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/)
- [Kubernetes Ingress Tutorial For Beginners](https://devopscube.com/kubernetes-ingress-tutorial/)
