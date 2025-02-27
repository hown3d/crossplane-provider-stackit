# Provider STACKIT

> [!WARNING]
> This is not a product officialy supported by STACKIT

`provider-stackit` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
STACKIT API.

## Getting Started

> [!NOTE]
> This provider depends on a [bug fix](https://github.com/crossplane/upjet/pull/448) in the upjet code for code generation.
> Please ensure that your go.work or go.mod uses a replace to point to the fixed version

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/stackitcloud/crossplane-provider-stackit):
```
up ctp provider install stackitcloud/crossplane-provider-stackit:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-stackit
spec:
  package: stackitcloud/crossplane-provider-stackit:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/stackitcloud/crossplane-provider-stackit).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/stackitcloud/crossplane-provider-stackit/issues).
