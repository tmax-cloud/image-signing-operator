package apis

import v1 "github.com/tmax-cloud/image-signing-operator/pkg/apiserver/apis/v1"

func init() {
	AddApiFuncs = append(AddApiFuncs, v1.AddV1Apis)
}
