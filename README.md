# :coin: Coinlog

[![Continuous Integration](https://github.com/maestre3d/coinlog/actions/workflows/ci.yml/badge.svg)](https://github.com/maestre3d/coinlog/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/maestre3d/coinlog)](https://goreportcard.com/report/github.com/maestre3d/coinlog)
[![codebeat badge](https://codebeat.co/badges/5093f302-8019-4ede-a7df-4eb8f816b8cb)](https://codebeat.co/projects/github-com-maestre3d-coinlog-master)
[![codecov](https://codecov.io/gh/maestre3d/coinlog/branch/master/graph/badge.svg?token=x2772NHOT7)](https://codecov.io/gh/maestre3d/coinlog)
[![Go Version](https://img.shields.io/github/go-mod/go-version/maestre3d/coinlog?style=square)](https://github.com/NeutrinoCorp/coinlog/blob/master/go.mod)

`Coinlog` is an assistant to keep track of your personal finance records.

- [:coin: Coinlog](#coin-coinlog)
  - [How-To](#how-to)
    - [Setup local environment](#setup-local-environment)
      - [Apache Kafka deployment](#apache-kafka-deployment)
      - [Postgres deployment](#postgres-deployment)
      - [Backend Application deployment](#backend-application-deployment)
      - [Frontend Web Application deployment](#frontend-web-application-deployment)

## How-To

### Setup local environment

`Coinlog` is ready to deploy its infrastructure using _Kubernetes (K8s)_ or _docker compose_. In this tutorial, 
_K8s_ is preferred.

First, create namespace and select it:

```shell
kubectl apply -f deployments/global.yml
```

```shell
kubectl config set-context --current --namespace=coinlog
```

#### Apache Kafka deployment

Build the Apache Kafka docker image contained in [kafka deployments folder](deployments/kafka).

```shell
docker build -t coinlog/kafka-kraft ./deployments/kafka
```

_This image is ready to use Kafka 3.4.0 with KRaft consensus protocol (Apache Zookeeper-less)._

Then, deploy the image:

```shell
kubectl apply -f deployments/kafka/kafka.yml
```

This will deploy 3 Kafka nodes available at `coinlog-kafka.coinlog.svc.cluster.internal` port `9092`.

#### Postgres deployment

The Postgres deployment requires to manually create the database `coinlog` after its deployment. DO NOT deploy
application if database was not created otherwise they will get stuck in **CrashLoop** state.

If `kubegres` operator not installed, please run the following command:

```shell
kubectl apply -f https://raw.githubusercontent.com/reactive-tech/kubegres/v1.16/kubegres.yaml
```

Reference [here](https://www.kubegres.io/doc/getting-started.html).

Then, perform the actual Postgres node deployments:

```shell
kubectl apply -f deployments/postgres/postgres.yml
```

_This will deploy 1 master node and 2 replicas._

Get access to shell in Postgres master node (_pod/coinlog-postgres-1-0_):

```shell
kubectl exec -it pod/coinlog-postgres-1-0 -- /bin/sh
```

Then run (use **root** as _postgres_ user password) to create database:

```shell
psql user=postgres
```

```shell
postgres=# CREATE DATABASE coinlog;
```

#### Backend Application deployment

First, build the docker image:

```shell
docker build -t coinlog/http-api:0.0.1 -f ./deployments/coinlog-http-api/Dockerfile .
```

_NOTE: Use image tags to perform rolling updates for deployments. Every change will require to update the K8s YAML._

If `nginx ingress controller` not installed, please run the following command:

```shell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.5.1/deploy/static/provider/cloud/deploy.yaml
```

Reference [here](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start).

Then, perform the actual deployment:

```shell
kubectl apply -f deployments/coinlog-http-api/coinlog.yml
```

_This will deploy 3 node stateless replicas._

If `api.coinlog.info` hostname was not set, run the command:

```shell
echo "127.0.0.1 api.coinlog.info" >> /etc/hosts
```

_This will enable external traffic to nginx edge proxy_.

Finally, export application nodes to cluster-external traffic:

```shell
kubectl port-forward services/coinlog-http-api-svc 8080:8080
```

_Or just use nginx ingress (localhost:80 or api.coinlog.info)._

More information about `k8s nginx ingress controller` [here](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/).

#### Frontend Web Application deployment

First, build the docker image:

```shell
docker build -t coinlog/web-client:0.0.1 -f ./deployments/coinlog-web-client/Dockerfile ./client/coinlog-web
```

_NOTE: Use image tags to perform rolling updates for deployments. Every change will require to update the K8s YAML._

Then, perform the actual deployment:

```shell
kubectl apply -f deployments/coinlog-web-client/coinlog-web.yml
```

_This will deploy 3 node stateless replicas._

If `app.coinlog.info` hostname was not set, run the command:

```shell
echo "127.0.0.1 api.coinlog.info" >> /etc/hosts
```

_This will enable external traffic to nginx edge proxy_.

Finally, export application nodes to cluster-external traffic:

```shell
kubectl port-forward services/coinlog-http-api-svc 8080:8080
```

_Or just use nginx ingress (localhost:80 or api.coinlog.info)._

More information about `k8s nginx ingress controller` [here](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/).
