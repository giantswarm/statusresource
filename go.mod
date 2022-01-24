module github.com/giantswarm/statusresource/v3

go 1.14

require (
	github.com/giantswarm/apiextensions/v3 v3.39.0
	github.com/giantswarm/backoff v1.0.0
	github.com/giantswarm/errors v0.3.0
	github.com/giantswarm/exporterkit v1.0.0
	github.com/giantswarm/k8sclient/v7 v7.0.1
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/micrologger v0.6.0
	github.com/giantswarm/operatorkit/v7 v7.0.0
	github.com/giantswarm/tenantcluster/v4 v4.1.0
	github.com/google/go-cmp v0.5.7
	github.com/prometheus/client_golang v1.12.0
	k8s.io/apimachinery v0.20.12
	k8s.io/client-go v0.20.12
)

replace sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
