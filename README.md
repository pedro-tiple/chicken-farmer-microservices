# Chicken Farmer - Micro-services
Portfolio project to demonstrate knowledge in a variety of modern technologies.

## How to run it:
Dependencies: Kubernetes (kubectl), Minikube, Docker

minikube start
kubectl apply -f https://getambassador.io/yaml/ambassador/ambassador-rbac.yaml
kubectl apply -f k8s/ambassador-svc.yaml
kubectl apply -f k8s/redis-svc.yaml
kubectl apply -f chicken-svc/chicken-svc.yaml
kubectl apply -f barn-svc/barn-svc.yaml
kubectl apply -f user-svc/user-svc.yaml
minikube service list

kubectl patch svc chicken-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc barn-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc user-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc redis-svc --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'

cd ../chicken-svc
docker build -t ptiple/chicken-svc:latest .
docker push ptiple/chicken-svc:latest
cd ../barn-svc
docker build -t ptiple/barn-svc:latest .
docker push ptiple/barn-svc:latest
cd ../user-svc
docker build -t ptiple/user-svc:latest .
docker push ptiple/user-svc:latest

kubectl replace --force -f chicken-svc.yaml
kubectl replace --force -f barn-svc.yaml
kubectl replace --force -f user-svc.yaml
kubectl replace --force -f user-svc.yaml

Connect to pod:
$ kubectl get pods
$ kubectl exec -it [PODNAME] bash