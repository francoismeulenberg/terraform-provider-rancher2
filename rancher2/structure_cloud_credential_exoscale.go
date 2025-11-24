package rancher2

// Flatteners

func flattenCloudCredentialExoscale(in *exoscaleCredentialConfig, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	if len(in.APIKey) > 0 {
		obj["api_key"] = in.APIKey
	}

	if len(in.APISecretKey) > 0 {
		obj["api_secret_key"] = in.APISecretKey
	}

	return []interface{}{obj}
}

// Expanders

func expandCloudCredentialExoscale(p []interface{}) *exoscaleCredentialConfig {
	obj := &exoscaleCredentialConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["api_key"].(string); ok && len(v) > 0 {
		obj.APIKey = v
	}

	if v, ok := in["api_secret_key"].(string); ok && len(v) > 0 {
		obj.APISecretKey = v
	}

	return obj
}
