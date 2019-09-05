# Chicken Farmer - Micro-services
Portfolio project to demonstrate knowledge in a variety of modern technologies.

## How to run it:
Dependencies: Kubernetes (kubectl), Minikube, Docker

kubectl apply -f https://getambassador.io/yaml/ambassador/ambassador-rbac.yaml
kubectl apply -f ambassador-svc.yaml
kubectl apply -f chicken-svc.yaml
minikube service list

docker build -t ptiple/chicken-svc:latest .
docker push ptiple/chicken-svc:latest
kubectl replace --force -f chicken-svc.yaml

kubectl patch svc chicken-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
