# Kuberiter

Kuberiter, write your articles and publish them to your Kubernetes cluster.

![cover](./docs/images/writer.jpg)

## Installation

- kubectl apply

    ```bash
    kubectl apply -f https://raw.githubusercontent.com/poneding/kuberiter/master/deploy/crds.yaml
    kubectl apply -f https://raw.githubusercontent.com/poneding/kuberiter/master/deploy/manifests.yaml
    ```

- helm install

    ```bash
    helm repo add kuberiter https://poneding.github.io/kuberiter
    helm install kuberiter kuberiter/kuberiter --namespace kuberiter --create-namespace
    ```

## Usage
