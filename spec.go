package statusresource

import providerv1alpha1 "github.com/giantswarm/apiextensions/v6/pkg/apis/provider/v1alpha1"

type Patch struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type Status struct {
	Cluster providerv1alpha1.StatusCluster `json:"cluster" yaml:"cluster"`
}

type Provider interface {
	ClusterStatus() providerv1alpha1.StatusCluster
}
