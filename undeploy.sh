#!/bin/bash

kubectl delete deployments --all
echo " - [*] deployments deleted"

kubectl delete services --all
echo " - [*] services deleted"

kubectl delete pods --all
echo " - [*] pods deleted"

kubectl delete cm eververse-config
echo " - [*] eververse-config cm deleted"