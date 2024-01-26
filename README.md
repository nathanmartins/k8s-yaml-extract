# k8s-yaml-extract

`k8s-yaml-extract`, is a powerful command-line interface (CLI) tool built for extraction of Kubernetes YAML manifests. This tool streamlines the process of filtering YAML files based on specified criteria, offering enhanced flexibility and control.

Key Features:

--name Flag: Extract YAMLs matching a specific resource name. ( REQUIRED )
--kind Flag: Extract YAMLs based on a particular Kubernetes resource kind.

Combined Filtering: Use both --name and --kind flags simultaneously for granular extraction.


## Examples

It can work just by piping in from STDIN:

`kustomize build example/ | k8s-yaml-extract --kind=deployment`

`kustomize build example/ | k8s-yaml-extract --name=my-app --kind=deployment`

or a static existing file as a parameter:

`k8s-yaml-extract --name=my-app --kind=deployment example/out.yaml`
