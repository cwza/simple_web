## Build and Run
``` sh
go build -o app
PORT=8080 ./app
curl http://localhost:8080/hello
```

## Deploy to Dockerhub
When you push to master branch the github action will automatically build image and push it to my dockerhub

## Deploy to k8s
``` sh
kubectl create namespace try
cd helm
# One container
helm install --set name=simple-web -n try -f values.yaml simple-web .
# Two container
# helm install --set name=simple-web --set multiple=true -n try -f values.yaml simple-web .
helm delete simple-web --namespace=try
```