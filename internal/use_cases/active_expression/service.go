package activeexpression

type Service struct {
	repo repo
}

func NewSvc(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Do() response {
	expressions := s.repo.GetExpression()
	response := response{
		Expressions: make([]string, len(expressions)),
	}
	for i := range expressions {
		response.Expressions[i] = expressions[i].Val
	}
	return response
}
