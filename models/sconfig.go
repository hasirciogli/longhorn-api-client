package models

type SConfig struct {
	LonghornUiEndpoint string
	BasicAuth          *string
}

func NewConfig(longhronUiEndpoint string, basicAuth *string) *SConfig {
	if basicAuth == nil {
		return &SConfig{
			LonghornUiEndpoint: longhronUiEndpoint,
		}
	}

	return &SConfig{
		LonghornUiEndpoint: longhronUiEndpoint,
		BasicAuth:          basicAuth,
	}
}
