# Chicken Farmer - Micro-services
Portfolio project to demonstrate knowledge in a variety of modern technologies.

## How to run it:
Dependencies: Kubernetes (kubectl), Minikube, Docker

minikube start
kubectl apply -f https://getambassador.io/yaml/ambassador/ambassador-rbac.yaml
kubectl apply -f k8s/ambassador-svc.yaml
kubectl apply -f k8s/redis-svc.yaml
kubectl apply -f k8s/frontend-svc.yaml
kubectl apply -f services/chicken-svc/chicken-svc.yaml
kubectl apply -f services/barn-svc/barn-svc.yaml
kubectl apply -f services/user-svc/user-svc.yaml
kubectl apply -f services/time-svc/time-svc.yaml
minikube service list

Open NodePorts for local development:
kubectl patch svc chicken-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc barn-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc user-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc time-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc redis-svc --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'

Build code and images:
cd services/chicken-svc
CGO_ENABLED=0 go build -o chicken-svc main/main.go
docker build -t ptiple/chicken-svc:latest .
docker push ptiple/chicken-svc:latest
cd ../barn-svc
CGO_ENABLED=0 go build -o barn-svc main/main.go
docker build -t ptiple/barn-svc:latest .
docker push ptiple/barn-svc:latest
cd ../user-svc
CGO_ENABLED=0 go build -o user-svc main/main.go
docker build -t ptiple/user-svc:latest .
docker push ptiple/user-svc:latest
cd ../time-svc
CGO_ENABLED=0 go build -o time-svc main/main.go
docker build -t ptiple/time-svc:latest .
docker push ptiple/time-svc:latest
cd ../../front-end
npm run build
docker build -t ptiple/chicken-farmer-microservices-frontend:latest .
docker push ptiple/chicken-farmer-microservices-frontend:latest

Reload configs and rebuild pods:
kubectl replace --force -f k8s/ambassador-svc.yaml
kubectl replace --force -f k8s/redis-svc.yaml
kubectl replace --force -f k8s/frontend-svc.yaml
kubectl replace --force -f services/chicken-svc/chicken-svc.yaml
kubectl replace --force -f services/barn-svc/barn-svc.yaml
kubectl replace --force -f services/user-svc/user-svc.yaml
kubectl replace --force -f services/time-svc/time-svc.yaml

Build image for whole project
docker build -t ptiple/chicken-farmer-microservices:latest .
docker push ptiple/chicken-farmer-microservices:latest

Connect to pod:
$ kubectl get pods
$ kubectl exec -it [PODNAME] bash