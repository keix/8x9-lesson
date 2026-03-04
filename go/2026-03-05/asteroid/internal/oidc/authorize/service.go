package authorize

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Authorize() (*Result, error) {
	return &Result{
		RedirectURL: "",
	}, nil
}
