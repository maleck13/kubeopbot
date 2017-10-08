#!/bin/bash

NAMESPACE=kubeop-bot
kubectl create ns ${NAMESPACE}
kubectl create -f ./artifacts/example/sa.yaml -n ${NAMESPACE}
kubectl create -f ./artifacts/example/auth-delegator.yaml -n kube-system
kubectl create -f ./artifacts/example/auth-reader.yaml -n kube-system
kubectl create -f ./artifacts/example/rc.yaml -n ${NAMESPACE}
kubectl create -f ./artifacts/example/service.yaml -n ${NAMESPACE}