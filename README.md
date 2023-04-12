# Kuberiter

## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

```bash
helm repo add kuberiter https://poneding.github.io/kuberiter
```

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
kuberiter` to see the charts.

To install the kuberiter chart:

```bash
helm install kuberiter kuberiter/kuberiter
```

To uninstall the chart:

```bash
helm uninstall kuberiter
```
