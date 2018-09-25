package bootkube

const (
	// MachineConfigOperator00ConfigCrd is the constant to represent contents of Machine_ConfigOperator00ConfigCrd.yaml file
	MachineConfigOperator00ConfigCrd = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: mcoconfigs.machineconfiguration.openshift.io
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Namespaced
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: mcoconfigs
    # singular name to be used as an alias on the CLI and for display
    singular: mcoconfig
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: MCOConfig
`
)
