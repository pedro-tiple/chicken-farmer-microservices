# Chicken Farmer - Micro-services - WIP
Portfolio project to demonstrate knowledge in a variety of modern technologies.

AKA: the over-engineered one

Note: Ongoing WIP, making changes as new things to try out come up.

## How to run it (outdated):
Dependencies: Kubernetes (kubectl), Minikube, Docker

minikube start
```
kubectl apply -f https://getambassador.io/yaml/ambassador/ambassador-rbac.yaml
kubectl apply -f k8s/ambassador-svc.yaml
kubectl apply -f k8s/redis-svc.yaml
kubectl apply -f k8s/frontend-svc.yaml
kubectl apply -f services/chicken-svc/chicken-svc.yaml
kubectl apply -f services/barnsvc/barnsvc.yaml
kubectl apply -f services/farmersvc/farmersvc.yaml
kubectl apply -f services/timesvc/timesvc.yaml
minikube service list
```

Open NodePorts for local development:
```
kubectl patch svc chicken-svc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc barnsvc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc farmersvc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc timesvc-mongodb --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
kubectl patch svc redis-svc --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"}]'
```

Build code and images:
```
cd services/chicken-svc
CGO_ENABLED=0 go build -o chicken-svc main/main.go
docker build -t ptiple/chicken-svc:latest .
docker push ptiple/chicken-svc:latest
cd ../barnsvc
CGO_ENABLED=0 go build -o barnsvc main/main.go
docker build -t ptiple/barnsvc:latest .
docker push ptiple/barnsvc:latest
cd ../farmersvc
CGO_ENABLED=0 go build -o farmersvc main/main.go
docker build -t ptiple/farmersvc:latest .
docker push ptiple/farmersvc:latest
cd ../timesvc
CGO_ENABLED=0 go build -o timesvc main/main.go
docker build -t ptiple/timesvc:latest .
docker push ptiple/timesvc:latest
cd ../../front-end
npm run build
docker build -t ptiple/chicken-farmer-microservices-frontend:latest .
docker push ptiple/chicken-farmer-microservices-frontend:latest
```

Reload configs and rebuild pods:
```
kubectl replace --force -f k8s/ambassador-svc.yaml
kubectl replace --force -f k8s/redis-svc.yaml
kubectl replace --force -f k8s/frontend-svc.yaml
kubectl replace --force -f services/chicken-svc/chicken-svc.yaml
kubectl replace --force -f services/barnsvc/barnsvc.yaml
kubectl replace --force -f services/farmersvc/farmersvc.yaml
kubectl replace --force -f services/timesvc/timesvc.yaml
```

Build image for whole project
```
docker build -t ptiple/chicken-farmer-microservices:latest .
docker push ptiple/chicken-farmer-microservices:latest
```

Connect to pod:
```
kubectl get pods
kubectl exec -it [PODNAME] bash
```