package calcexpression

import (
	"strings"

	"prettyprj/internal/entities"
	appErrors "prettyprj/internal/errors"
)

type Service struct {
	repo repo
}

func NewSvc(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Do(request *request) error {

	if strings.Contains(request.Expression, "&") {
		return appErrors.ErrUnsupportedOperation
	}

	expression := entities.Expression{
		Val: request.Expression,
	}
	s.repo.SaveExpression(expression)
	return nil
}
