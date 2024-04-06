package calcexpression

import "prettyprj/internal/entities"

type repo interface {
	SaveExpression(expression entities.Expression)
}
