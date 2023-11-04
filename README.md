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

## configs

For Sanjab server configs, you need to create a ```config.yml``` file. This file's structure
is like this:

```yaml
timeout: 60 # in seconds
port: 80 # service http port
namespace: "default" # kubernetes namespace
ceph_disable: false # enabling ceph upload
ceph:
  host: "http://127.0.0.1:6800"
  access: "access-token"
  secret: "secret-token"
  bucket: "bucket-name"
```

You can list your desire resources in the ```config.yml``` file. Sanjab creates a worker process for each
resource to manage them concurrently. For example, if we want to back up only pods and deployments, we need
to set our resources like this:

```yaml
resources:
  - pods
  - deployments
```
