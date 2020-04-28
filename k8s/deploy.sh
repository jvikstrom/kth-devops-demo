#!/bin/bash
kubectl apply -f server.yml
kubectl apply -f client.yml
kubectl apply -f prometheus-config.yml
kubectl apply -f prometheus.yml
