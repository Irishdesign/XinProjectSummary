package controllers

import (
	"XinAPI/app/http/models"
	"XinAPI/pkg/l"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ServeStatic(c *gin.Context) {

	m := new(models.EcpayCashFlow)
	// need different orderNo and to refresh everytime you want to test order submit
	orderNo := "O2020032600046" // revise here
	m.GetEcpayParams(orderNo)
	c.HTML(http.StatusOK, "ecpay_submit.html", gin.H{
		"TotalAmount":       m.TotalAmount,
		"TradeDesc":         m.TradeDesc,
		"ItemName":          m.ItemName,
		"ReturnURL":         m.ReturnURL,
		"MerchantTradeNo":   m.MerchantTradeNo,
		"MerchantTradeDate": m.MerchantTradeDate,
		"ChoosePayment":     m.ChoosePayment,
		"EncryptType":       m.EncryptType,
		"MerchantID":        m.MerchantID,
		"CheckMacValue":     m.CheckMacValue,
		"PaymentType":       m.PaymentType,
	})
}

type EcpayRes struct {
	CustomField1         string `json:"CustomField1" db:"CustomField1"`
	CustomField2         string `json:"CustomField2" db:"CustomField2"`
	CustomField3         string `json:"CustomField3" db:"CustomField3"`
	CustomField4         string `json:"ReturnURL" db:"ReturnURL"`
	MerchantID           string `json:"MerchantID" db:"MerchantID"`
	MerchantTradeNo      string `json:"MerchantTradeNo" db:"order_id"`
	PaymentDate          string `json:"PaymentDate" db:"created_at"`
	PaymentType          string `json:"PaymentType" db:"PaymentType"`
	PaymentTypeChargeFee string `json:"PaymentTypeChargeFee" db:"PaymentTypeChargeFee"`
	RtnCode              string `json:"RtnCode" db:"status"`
	RtnMsg               string `json:"RtnMsg" db:"RtnMsg"`
	SimulatePaid         string `json:"SimulatePaid" db:"simulation"`
	StoreID              string `json:"StoreID" db:"StoreID"`
	TradeAmt             string `json:"TradeAmt" db:"TradeAmt"`
	TradeDate            string `json:"TradeDate" db:"TradeDate"`
	TradeNo              string `json:"TradeNo" db:"TradeNo"`
	CheckMacValue        string `json:"CheckMacValue" db:"CheckMacValue"`
}

func GetFromEcpay(c *gin.Context) {
	db, _ := models.NewDBSolomo()
	data := &EcpayRes{}
	// get res data by defined struct
	c.Bind(data)

	d := EcpayRes{
		MerchantID:           data.MerchantID,
		MerchantTradeNo:      data.MerchantTradeNo,
		PaymentDate:          data.PaymentDate,
		PaymentType:          data.PaymentType,
		PaymentTypeChargeFee: data.PaymentTypeChargeFee,
		RtnCode:              data.RtnCode,
		RtnMsg:               data.RtnMsg,
		SimulatePaid:         data.SimulatePaid,
		TradeAmt:             data.TradeAmt,
		TradeDate:            data.TradeDate,
		TradeNo:              data.TradeNo,
		CheckMacValue:        data.CheckMacValue,
	}

	// get OrderID
	o := new(models.Order)
	// err := o.FindByNo(d.MerchantTradeNo)
	err := o.FindByNo(data.MerchantTradeNo)
	if err != nil {
		l.Log(`***save failed ******`, err)
		return
	}
	time, _ := time.Parse("2020/04/15 11:57:25", d.PaymentDate)
	// Struct to JSON
	b, _ := json.Marshal(d)

	res := models.Cashflow{
		OrderID:    o.ID,
		Simulation: convertToInt(d.SimulatePaid),
		Status:     convertToInt(d.RtnCode),
		CreatedAt:  time,
		Payload:    b,
	}

	err = db.Save(&res)

	if err != nil {
		l.Log(`***save err ******`, err)
	}
	l.Log(`***save to DB: cashflows ******`, res)
}

func convertToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
