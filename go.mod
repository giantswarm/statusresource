module github.com/giantswarm/statusresource/v3

go 1.14

require (
	github.com/giantswarm/apiextensions/v3 v3.39.0
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/errors v0.3.0
	github.com/giantswarm/exporterkit v0.2.0
	github.com/giantswarm/k8sclient/v5 v5.11.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/giantswarm/operatorkit/v4 v4.0.0
	github.com/giantswarm/tenantcluster/v4 v4.0.0
	github.com/google/go-cmp v0.5.6
	github.com/prometheus/client_golang v1.8.0
	k8s.io/apimachinery v0.18.19
	k8s.io/client-go v0.18.19
)

replace sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
