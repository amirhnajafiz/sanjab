package enum

import "fmt"

type Resource string

const (
	PodResource            Resource = "pods"
	DeploymentResource     Resource = "deployments"
	ServiceResource        Resource = "services"
	CronjobResource        Resource = "cronjobs"
	ConfigMapResource      Resource = "configmaps"
	SecretResource         Resource = "secrets"
	ServiceAccountResource Resource = "serviceaccounts"
	StatefulResource       Resource = "statefulsets"
	HPAResource            Resource = "hpas"
	IngressResource        Resource = "ingresses"
	PVCResource            Resource = "pvcs"
)

func (r Resource) ToString() string {
	return fmt.Sprintf("%s", r)
}
