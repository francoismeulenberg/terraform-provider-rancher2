package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Types

type exoscaleCredentialConfig struct {
	APIKey       string `json:"ApiKey,omitempty" yaml:"ApiKey,omitempty"`
	APISecretKey string `json:"ApiSecretKey,omitempty" yaml:"ApiSecretKey,omitempty"`
}

//Schemas

func cloudCredentialExoscaleFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"api_key": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Exoscale API Key",
		},
		"api_secret_key": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Exoscale API Secret Key",
		},
	}

	return s
}
