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
helm install --namespace=try -f values.yaml simple-web .
```