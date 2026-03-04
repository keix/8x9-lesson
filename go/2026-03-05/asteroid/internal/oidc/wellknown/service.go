package wellknown

type Service struct {
	issuer string
}

func NewService(issuer string) *Service {
	return &Service{
		issuer: issuer,
	}
}

func (s *Service) GetConfiguration() *Configuration {
	return &Configuration{
		Issuer:                           s.issuer,
		AuthorizationEndpoint:            s.issuer + "/authorize",
		TokenEndpoint:                    s.issuer + "/token",
		JwksURI:                          s.issuer + "/.well-known/jwks.json",
		ResponseTypesSupported:           []string{"code"},
		SubjectTypesSupported:            []string{"public"},
		IDTokenSigningAlgValuesSupported: []string{"RS256"},
	}
}
