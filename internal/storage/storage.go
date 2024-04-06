package storage

import (
	"sync"

	"prettyprj/internal/entities"
)

type storage struct {
	expressions []entities.Expression
	mx          *sync.Mutex
}

func New() *storage {
	return &storage{
		mx: &sync.Mutex{},
	}
}

func (s *storage) GetExpression() []entities.Expression {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.expressions
}

func (s *storage) SaveExpression(expression entities.Expression) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.expressions = append(s.expressions, expression)
}
