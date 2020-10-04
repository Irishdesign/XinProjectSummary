// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	"XinAPI/app/http/models"
	models_gen "XinAPI/build/gqlgen/api/models"
	"context"
)

func (r *queryResolver) EcpayCashflow(ctx context.Context, orderNo string) (*models_gen.REcpayCashflows, error) {
	// panic(fmt.Errorf("not implemented"))

	m := new(models.EcpayCashFlow)

	m.GetEcpayParams(orderNo)

	// url.QueryEscape()
	res := &models_gen.REcpayCashflows{
		Code: 1,
		Msg:  "SUCCESS",
		Data: m,
	}
	return res, nil
	// panic(fmt.Errorf("not implemented"))
}
