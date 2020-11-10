package trust

import "github.com/tmax-cloud/image-signing-operator/internal/utils"

const (
	DctEnvKeyRoot   = "DOCKER_CONTENT_TRUST_ROOT_PASSPHRASE"
	DctEnvKeyTarget = "DOCKER_CONTENT_TRUST_REPOSITORY_PASSPHRASE"
	TrustRoleRoot   = RoleType("root")
	TrustRoleTarget = RoleType("target")
)

type RoleType string

var RoleMap = map[RoleType]string{
	TrustRoleRoot:   DctEnvKeyRoot,
	TrustRoleTarget: DctEnvKeyTarget,
}

type TrustPass map[string]string

func NewTrustPass() TrustPass {
	pass := make(TrustPass)
	pass[DctEnvKeyRoot] = utils.RandomString(12)
	pass[DctEnvKeyTarget] = utils.RandomString(12)

	return pass
}

func (t TrustPass) AssignNewTargetPass() {
	t[DctEnvKeyTarget] = utils.RandomString(12)
}
