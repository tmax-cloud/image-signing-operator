module github.com/tmax-cloud/image-signing-operator

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/gorilla/mux v1.8.0
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/operator-framework/operator-lib v0.1.0
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.18.8
	k8s.io/kube-aggregator v0.18.8
	knative.dev/pkg v0.0.0-20201110224859-b713a3c08e6c
	sigs.k8s.io/controller-runtime v0.6.2
)

replace (
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/api => k8s.io/api v0.18.8
)
