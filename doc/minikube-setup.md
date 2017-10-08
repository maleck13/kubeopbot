download minikube

set the following values in the rc

```
- --requestheader-client-ca-file=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
- --requestheader-username-headers=X-Remote-User
- --requestheader-group-headers=X-Remote-Group
- --requestheader-extra-headers-prefix=X-Remote-Extra        
```