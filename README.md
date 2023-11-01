# Sanjab

Sanjab (aka squirrel is english) is a service for backing up your k8s objects into Ceph cluster. Whenever a new request is
sent to api-server, sanjab gets the kubernetes object of that request and stores it Ceph cluster.

Sanjab is built with kubernetes operator pattern. It watches over the following resources on cluster, and stores the object
as a yaml file on Ceph cluster.

- Pods
- Deployments
- Services
- Cronjobs
- Configmaps
- Secrets
- ServiceAccounts
- StatefulSets
- HPAs
- Ingress
- PVCs
