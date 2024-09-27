# Session 1: Kubernetes Basics

In this session, we will learn what is Kubernetes and why we need it. We will learn about different components of Kubernetes. On hands-on section, we will learn to create Kubernetes cluster and interact with it. 

**Theory:**
- Introduction to Kubernetes
- Kubernetes Architecture and components

**Hands On:**
- Create a Kubernetes cluster using or `kind`
- Interacting with Kubernetes cluster using `kubectl`
- Interacting with Kubernetes cluster using UI (Kubernetes Dashboard, K9s etc.)

## Requirements
- Install [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- Install [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

## Theory Resources
- [Kubernetes Concepts (Official docs)](https://kubernetes.io/docs/concepts/)
- [Why do you need Kubernetes? What Kubernetes is not?](https://www.linkedin.com/pulse/why-do-you-need-kubernetes-krishna-wattamwar-vs8of)
- [Why you should use Kubernetes?](https://faun.pub/why-you-should-use-kubernetes-bf395bef52de)
- [7 Reasons Kubernetes Is Important for DevOps](https://www.turing.com/blog/importance-of-kubernetes-for-devops)
- [Understanding Kubernetes Architecture: A Comprehensive Guide](https://devopscube.com/kubernetes-architecture-explained/)
## Creating a Kubernetes Cluster

We will be creating a 3 node cluster. At first create a `kind.yaml` file with the following content:

```yaml
# three node (two workers) cluster config
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
```

Now, run the following command to create the cluster:

```bash
kind create cluster --name=my-first-k8s-cluster --config=kind.yaml
```

To view all the created clusters on your machine, run:
```bash
kind get clusters
```


## Interacting with Kubernetes cluster using `kubectl`

To view kubeconfig, run:
```bash
kubectl config view

```

```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://127.0.0.1:45573
  name: kind-my-first-k8s-cluster
contexts:
- context:
    cluster: kind-my-first-k8s-cluster
    user: kind-my-first-k8s-cluster
  name: kind-my-first-k8s-cluster
current-context: kind-my-first-k8s-cluster
kind: Config
preferences: {}
users:
- name: kind-my-first-k8s-cluster
  user:
    client-certificate-data: DATA+OMITTED
    client-key-data: DATA+OMITTED
```

To view sensitive information in the kubeconfig, run:
```bash
kubectl config view --raw
```

To view information regarding your cluster, run:
```bash
kubectl cluster-info --context kind-my-first-k8s-cluster
```

  To view list of nodes in your cluster, run:
```bash
kubectl get nodes  
```

To view list of namespaces in your cluster, run:
```bash
kubectl get namespaces
```

To view all pods in the cluster, run:
  
```bash
kubectl get pods --all-namespaces

# you can use shorthand -A instead of --all-namespaces
kubectl get pods -A
```

To view pods of a particular namespace, run:

```bash
kubectl get pods -n kube-system
```

If you don't specify namespace, it will show pods from `default` namespace:
```bash
kubectl get pods
```

To view details of a pod, run:

```bash
kubectl describe pod -n <namespace> <pod name>

# example
kubectl describe pod -n kube-system kube-apiserver-my-first-k8s-cluster-control-plane 
```

To view YAML definition of a pod, run:

```bash
kubectl get pod -n <namespace> <pod name> -o yaml

# example
kubectl get pod -n kube-system kube-apiserver-my-first-k8s-cluster-control-plane -o yaml
```

To view logs of a pod, run:
```bash
kubectl logs -n <namespace> <pod name>

# example
kubectl logs -n kube-system kube-apiserver-my-first-k8s-cluster-control-plane
```

To create a resource from a YAML manifest, run:
```bash
kubectl create -f manifests/pod.yaml
```

Apply changes made on a resource manifest, run following command. This will also create the resource if it does not exist.
```bash
kubectl apply -f manifests/pod.yaml
```

To lunch a shell inside a pod, run:
```bash
kubectl exec -it busybox -n default -- /bin/sh 
```

To delete a pod, run:
```bash
kubectl delete pod -n default busybox

# you can also use manifest file
kubectl delete -f manifests/pod.yaml
```

To view available commands in `kubectl`, run:
```bash
kubectl --help
```


**Resources:**

- [kubectl reference guide](https://kubernetes.io/docs/reference/kubectl/quick-reference/)
- [Commands guide for kubectl](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands)


## Interacting with Kubernetes cluster with UI

### Using k9s Terminal UI
- Install [k9s](https://k9scli.io)https://k9scli.io
- Run `k9s`

Press `:` to enter into command mode and `/` into search mode.

### Using Kubernetes Dashboard

- Install [Helm](https://helm.sh/docs/intro/install/)
- Install Kuberentes dashboard
  
```bash
# Add kubernetes-dashboard repository
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/

# Deploy a Helm Release named "kubernetes-dashboard" using the kubernetes-dashboard chart
helm upgrade --install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --create-namespace --namespace kubernetes-dashboard
```

- Port-forward dashboard Service

```bash
kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-kong-proxy 8443:443
```

- Now, browse https://localhost:8443/ . It will ask for a token.

- Create an user for the dashboard
```bash
kubectl apply -f manifests/dashboard-user.yaml
```

- Create a token for the user
  
```bash
kubectl -n kubernetes-dashboard create token admin-user
```

- Use the printed token to login into the dashboard


## Delete cluster
To delete the cluster and everything in it, run:

```bash
kind delete cluster --name=my-first-k8s-cluster
```
