# Session 2: Deploying an Application on Kubernetes

In this session, we will learn how to deploy an application on Kubernetes. We will learn about how we can use different types of workloads to deploy our app based on different requirements.

**Theory:**

- Introduction to Pod, Deployment, ReplicaSet
- Introduction to Job, CronJob
- Introduction to DaemonSet

**Hands On:**

- Deploy a sample application using Pod
- Deploy same application using Deployment
  - Scale up a deployment
  - Perform a rolling update
  - Perform a rollback
- Deploy a job
- Schedule a CronJob
- Deploy a DaemonSet

## Deploy Sample Application using Pod

To deploy the sample web server using a Pod, run:

```bash
kubectl apply -f manifests/pod.yaml
```

To check, if the Pod is running, run:

```bash
kubectl get pod sample-web-server --show-labels
```

You can also run following commands to show the node where the Pod is running:

```bash
kubectl get pod sample-web-server -o wide
```

To access the web server, run this command on a separate terminal:

```bash
kubectl port-forward sample-web-server 8080:8080
```

Now, you can `curl` the web server using:

```bash
curl http://localhost:8080/hello
```

To delete the Pod, run:

```bash
kubectl delete pod sample-web-server
```

**Resources:**

- [Pod Documentation](https://kubernetes.io/docs/concepts/workloads/pods/)
- [What is Kubernetes Pod? Explained With Practical Examples](https://devopscube.com/kubernetes-pod/)
- [Kubernetes Pod Lifecycle Explained With Examples](https://devopscube.com/kubernetes-pod-lifecycle/)
- [Kubernetes Init Containers: A Complete Guide](https://devopscube.com/kubernetes-init-containers/)

## Deploy Sample Application using Deployment

To deploy the sample web server using a Deployment, run:

```bash
kubectl apply -f manifests/deployment.yaml
```

To check, if the Deployment is ready, run:

```bash
kubectl get deployment sample-web-server
```

You should see `1/1` replica is ready on `READY` column.

Deployment will create a ReplicaSet and which will create the Pod. To list the ReplicaSet created by the Deployment, run:

```bash
kubectl get replicaset -l app=hello-server
```

To get the pods created by the ReplicaSet, run:

```bash
kubectl get pods -l app=hello-server -o wide
```

### Scale up a Deployment

Now, will run 3 instance of our application so that it can handle more load. To scale up the Deployment, run:

```bash
kubectl scale deployment sample-web-server --replicas=3
```

To check, if the Deployment is scaled up, run:

```bash
kubectl get pods -l app=hello-server -o wide
```

You should see 3 pods are running.

### Perform a Rolling Update

At first, let's modify our sample web-server and build a new `v2` image version with the change.

Once the image has been pushed to the registry, update the Deployment with the new image:

```bash
kubectl set image deployment/sample-web-server webserver=hossainemruz/sample-web-server:v2
```

To see rollout progress, run:

```bash
kubectl rollout status deployment/sample-web-server
```

This rolling update will create a new ReplicaSet with the new image and scale down the old ReplicaSet. To check the current ReplicaSets, run:

```bash
kubectl get replicaset -l app=hello-server -o wide
```

You should see 2 ReplicaSets. One with the old image and another with the new image. The old ReplicaSet will be scaled down to 0 once the new ReplicaSet is ready.

To confirm the new image is running, run:

```bash
kubectl logs -l app=hello-server
```

It will print log for all the hello-server pods. You should see the log reporting version to be `v2.0.0`.

### Perform a Rollback

At first, let's modify our sample web-server and introduce a bug which will cause the app to crash. Then, build a new `v3` image version with the bug.

Once the image has been pushed to the registry, update the Deployment with the new image:

```bash
kubectl set image deployment/sample-web-server webserver=hossainemruz/sample-web-server:v3
```

This time, the new pod will fail to start due to the bug. To check the status of the pods, run:

```bash
kubectl get pods -l app=hello-server -o wide
```

You should see the new pod is in `CrashLoopBackOff` state.

To check the rollout status, run:

```bash
kubectl rollout status deployment/sample-web-server
```

You should see only one replica has been updated with new image.

If you check the ReplicaSets, you will see it created only 1 replica of the new image and old ReplicaSet is still running with 3 replicas.

To rollback to the previous version, run:

```bash
kubectl rollout undo deployment/sample-web-server
```

Once the rollback is complete, if you check the ReplicaSet you will see the ReplicaSet with `v2` image has been scaled to 3 replicas and ReplicaSet with `v3` image has been scaled to 0 replica.

**Resources:**

- [Deployment Documentation](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

## Deploy a Job

To deploy a Job, run:

```bash
kubectl apply -f manifests/job.yaml
```

To see the Job status, run:

```bash
kubectl get job pi
```

To see the pod created by the Job, run:

```bash
kubectl get pods -l batch.kubernetes.io/job-name=pi
```

To delete the Job, run:

```bash
kubectl delete job pi
```

**Resources:**

- [Job Documentation](https://kubernetes.io/docs/concepts/workloads/controllers/job/)

## Schedule a CronJob

To deploy a CronJob, run:

```bash
kubectl apply -f manifests/cronjob.yaml
```

To see the CronJob status, run:

```bash
kubectl get cronjob sample-cronjob
```

To see the Job created by the CronJob, run:

```bash
kubectl get jobs -l app=sample-cronjob
```

To see the pods created by the Jobs, run:

```bash
kubectl get pods -l app=sample-cronjob
```

To delete the CronJob, run:

```bash
kubectl delete cronjob sample-cronjob
```

**Resources:**

- [CronJob Documentation](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/)


## Deploy a DaemonSet

To create a sample DaemonSet, run:

```bash
kubectl apply -f manifests/daemonset.yaml
```

To see the DaemonSet status, run:

```bash
kubectl get daemonset fluentd
```

To see the pods created by the DaemonSet, run:

```bash
kubectl get pods -l app=fluentd -o wide
```

You will see it has created a pod on each worker node node.

To delete the DaemonSet, run:

```bash
kubectl delete daemonset fluentd
```

**Resources:**

- [DaemonSet Documentation](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)
- [Kubernetes Daemonset: A Comprehensive Guide](https://devopscube.com/kubernetes-daemonset/)
