# k8s-yaml-extract

`k8s-yaml-extract`, is a powerful command-line interface (CLI) tool built for extraction of Kubernetes YAML manifests. This tool streamlines the process of filtering YAML files based on specified criteria, offering enhanced flexibility and control.

## Usage

Flags:

- `--name`: Extract YAMLs matching a specific resource name. ( REQUIRED )
- `--kind`: Extract YAMLs based on a particular Kubernetes resource kind.

Combined Filtering: Use both `--name` and `--kind` flags simultaneously for even more granular extraction.

## Examples

It can work just by piping in from STDIN:

`kustomize build example/ | k8s-yaml-extract --kind=deployment`

`kustomize build example/ | k8s-yaml-extract --name=my-app --kind=deployment`

or a static existing file as a parameter:

`k8s-yaml-extract --name=my-app --kind=deployment example/out.yaml`

## Installation

In order to use the `k8s-yaml-extract` command-line tool, you need to have Go (version 1.21) installed on your system.

`go install github.com/nathanmartins/k8s-yaml-extract@latest`

or download the binary over at:

https://github.com/nathanmartins/k8s-yaml-extract/releases/latest
