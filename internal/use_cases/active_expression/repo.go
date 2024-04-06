package activeexpression

import "prettyprj/internal/entities"

type repo interface {
	GetExpression() []entities.Expression
}
