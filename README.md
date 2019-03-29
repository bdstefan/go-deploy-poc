# go-deploy-poc
A demo app developed with Golang and deployed with Docker + Kubernetes 

## How to:
1. docker build -t go-deploy-poc .
2. docker stop gopoc
3. docker rm gopoc
4. docker run --name gopoc -d -p 3030:3030 go-deploy-poc