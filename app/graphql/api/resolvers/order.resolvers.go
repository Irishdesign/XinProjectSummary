// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	"XinAPI/app/http/models"
	generated "XinAPI/build/gqlgen/api"
	models_gen "XinAPI/build/gqlgen/api/models"
	"XinAPI/pkg/l"
	"context"
	"fmt"
)

func (r *mutationResolver) AddOrder(ctx context.Context, newOrder *models_gen.NewOrder) (*models_gen.REcpayCashflows, error) {
	newCreateCode := 0
	memberInfo := new(models.Member)
	memberInfo.FindByNo(newOrder.MemberNo)
	// save to table : orders
	o := models.Order{
		Status:        newCreateCode,
		Platform:      newOrder.Platform,
		Type:          newOrder.Type,
		MemberID:      memberInfo.ID,
        ...
		PaymentType:   newOrder.PaymentType,
		PaymentStatus: newCreateCode,
		InvoiceType:   *newOrder.InvoiceType,
	}
	err := o.Save()
	if err != nil {
		l.Log("save new order error", err)
	}
	// save to table : order_items
	cart := new(models.CartAndProducts)
	cart.FindByMidAndType(memberInfo.ID, newOrder.Type)
	var orderVendorArr []int
	for _, c := range *cart {
		item := new(models.OrderItem)
		item.OrderVendorID = c.VendorID
		item.OrderID = o.ID
	    ...
		item.Price = c.SalePrice
		// item.Discount =
		item.ShipmentStatus = 0
		if models.IsNonRepeat(orderVendorArr, c.VendorID) { // venderID
			orderVendorArr = append(orderVendorArr, c.VendorID)
		}
		err = item.Save()
	}
	// save to table : order_vendors
	for _, existV := range orderVendorArr {
		OrderVendor := new(models.OrderVendor)
	    ...
		err = OrderVendor.Save()
	}
	//get Cashflow parms
	e := new(models.EcpayCashFlow)
	e.GetEcpayParams(o.No)

	res := &models_gen.REcpayCashflows{
		Code: 1,
		Msg:  "SUCCESS",
		Data: e,
	}
	return res, nil
}



func (r *Resolver) Order() generated.OrderResolver             { return &orderResolver{r} }
func (r *Resolver) OrderItem() generated.OrderItemResolver     { return &orderItemResolver{r} }
func (r *Resolver) OrderVendor() generated.OrderVendorResolver { return &orderVendorResolver{r} }

type orderResolver struct{ *Resolver }
type orderItemResolver struct{ *Resolver }
type orderVendorResolver struct{ *Resolver }
