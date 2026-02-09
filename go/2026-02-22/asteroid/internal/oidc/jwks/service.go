package jwks

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetKeySet() *KeySet {
	return &KeySet{
		Keys: []Key{},
	}
}
