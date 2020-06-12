# hostnameServer on Kubernetes Cluster

Access to the root of a webserver to get the name and node of a webserver. 

## On Kubernetes Cluster

In a cluster deploy multiple replicas of the webserver. If you perform multiple GET requests you
should see different pod's name. The .yaml contains both the deployment and the service of the application:
`kubectl apply -f ./hostname-depl-and-svc.yaml`

## Give permissions to the pods

By default the pods doesn't have permissions required for the execution of the app. Check by yourself by looking at the logs: 
`kubectl logs -f -lapp=appLabel --all-containers=true`
To give the permissions create a role binding:
`kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default` 

## Http GET from a host in the same network

Perform some GET requests. Actually the browsers change the output few times. It's better
to use:
`curl nodeIP:servicePort`
See pods and nodes change time in time. It's indifferent wich ip node to use because the service will load balancing the request among all the pods and nodes.

## Clean Up

Delete the deployment and the service:
```
kubectl delete deployment hostname-deployment
kubectl delete service hostname
```


