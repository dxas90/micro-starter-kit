# Testing k8s deployment

```bash
# service account
kubectl create -f deploy/basic/micro-service-account.yaml
kubectl get sa
kubectl get clusterroles | grep micro-role
kubectl get ClusterRoleBinding | grep micro-role-binding
kubectl delete -f deploy/basic/micro-service-account.yaml

# account-srv
kubectl create -f deploy/basic/account.yaml
kubectl create -f deploy/basic/account-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=account-srv -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/basic/account.yaml
kubectl delete -f deploy/basic/account-svc.yaml

# Gateway
kubectl create -f deploy/basic/gateway.yaml
kubectl create -f deploy/basic/gateway-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=gateway -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f

kubectl delete -f deploy/basic/gateway.yaml
kubectl delete -f deploy/basic/gateway-svc.yaml
```