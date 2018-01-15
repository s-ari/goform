echo -n "<ROOT_PASSWORD>" > /tmp/root-password
kubectl create secret generic root-password --from-file=/tmp/root-password
echo -n "<USER_PASSWORD" > /tmp/user-password
kubectl create secret generic user-password --from-file=/tmp/user-password
kubectl get secrets
kubectl describe secrets/root-password
kubectl describe secrets/user-password

kubectl create -f deployment.yaml
kubectl get pods -w
kubectl create -f service.yaml
kubectl get svc -w
kubectl create -f job.yaml
kubectl get job -w
kubectl create -f ingress.yaml
kubectl get ing -w

kubectl delete -f job.yaml
kubectl delete -f ingress.yaml
kubectl delete -f service.yaml
kubectl delete -f deployment.yaml
delete secrets root-password
delete secrets user-password
