module github.com/martin-helmich/SampleOperator

go 1.13

require (
	github.com/operator-framework/operator-sdk v0.18.0
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	k8s.io/client-go => k8s.io/client-go v0.18.2 // Required by prometheus-operator
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
)
