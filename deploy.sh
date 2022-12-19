#!/bin/bash

cd k8s || exit
kubectl apply -f cm.yaml
echo " - [X] eververse-config created"

kubectl apply -f eververse-deployment.yaml
echo " - [X] eververse-deployment applied"

kubectl apply -f eververse-service.yaml
echo " - [X] eververse-service applied"

kubectl apply -f pv.yaml
echo " - [X] pv applied"

kubectl apply -f pvc.yaml
echo " - [X] pvc applied"

kubectl apply -f redis-deployment.yaml
echo " - [X] redis-deployment applied"

kubectl apply -f redis-service.yaml
echo " - [X] redis-deployment applied"

minikube dashboard