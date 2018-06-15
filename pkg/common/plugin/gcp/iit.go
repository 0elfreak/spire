package gcp

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type IdentityToken struct {
	jwt.StandardClaims

	AuthorizedParty string `json:"azp"`
	Google          Google `json:"google"`
}

type Google struct {
	ComputeEngine ComputeEngine `json:"compute_engine"`
}

type ComputeEngine struct {
	ProjectID                 string `json:"project_id"`
	ProjectNumber             int64  `json:"project_number"`
	Zone                      string `json:"zone"`
	InstanceID                string `json:"instance_id"`
	InstanceName              string `json:"instance_name"`
	InstanceCreationTimestamp int64  `json:"instance_creation_timestamp"`
}

func AttestationStepError(step string, cause error) error {
	return fmt.Errorf("Attempted GCP IID attestation but an error occured %s: %s", step, cause)
}
