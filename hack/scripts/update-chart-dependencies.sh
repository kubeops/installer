#!/bin/sh
set -e

helm repo add appscode https://charts.appscode.com/stable/ || true

helm dependency update charts/kube-ui-server
