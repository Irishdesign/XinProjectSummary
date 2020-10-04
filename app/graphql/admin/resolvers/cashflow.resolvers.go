// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	"XinAPI/app/http/models"
	generated "XinAPI/build/gqlgen/admin"
	models_gen "XinAPI/build/gqlgen/admin/models"
	"context"
	"fmt"
)

func (r *queryResolver) Cashflow(ctx context.Context, orderID int) (*models_gen.RCashflow, error) {
	m := new(models.Cashflow)
	m.FindByOid(orderID)
	res := &models_gen.RCashflow{
		Code: 1,
		Msg:  "SUCCESS",
		Data: m,
	}
	return res, nil
}


func (r *Resolver) Cashflow() generated.CashflowResolver { return &cashflowResolver{r} }

type cashflowResolver struct{ *Resolver }
