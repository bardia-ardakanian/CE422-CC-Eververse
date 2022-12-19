#!/bin/bash

kubectl run webkit --image=bardiaardakanian/webkit -i --tty -- sh

# Hello World
kubectl exec webkit -- curl -v -Ss http://eververse-service.default.svc.cluster.local:8080
# Get BTC
# Get BTC Cache
kubectl exec webkit -- curl -X POST -v -Ss http://eververse-service.default.svc.cluster.local:8080/get?name=ETH
kubectl exec webkit -- curl -X POST -v -Ss http://eververse-service.default.svc.cluster.local:8080/get?name=ETH
# Get Doge
kubectl exec webkit -- curl -X POST -v -Ss http://eververse-service.default.svc.cluster.local:8080/get?name=DOG