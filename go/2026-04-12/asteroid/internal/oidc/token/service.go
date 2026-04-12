package token

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ExchangeToken() (*Result, error) {
	return &Result{
		AccessToken: "",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}, nil
}
