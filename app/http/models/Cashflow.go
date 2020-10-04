package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

// Cashflow is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Cashflow struct {
	ID         int       `json:"id" db:"id"`
	OrderID    int       `json:"order_id" db:"order_id"`
	Simulation int       `json:"simulation" db:"simulation"`
	Status     int       `json:"status" db:"status"`
	Payload    []byte    `json:"payload" db:"payload"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type EcpayCashFlow struct {
	TotalAmount       string `json:"TotalAmount" db:"TotalAmount"`
	TradeDesc         string `json:"TradeDesc" db:"TradeDesc"`
	ItemName          string `json:"ItemName" db:"ItemName"`
	ReturnURL         string `json:"ReturnURL" db:"ReturnURL"`
	MerchantTradeNo   string `json:"MerchantTradeNo" db:"MerchantTradeNo"`
	MerchantTradeDate string `json:"MerchantTradeDate" db:"MerchantTradeDate"`
	PaymentType       string `json:"PaymentType" db:"PaymentType"`
	ChoosePayment     string `json:"ChoosePayment" db:"ChoosePayment"`
	EncryptType       string `json:"EncryptType" db:"EncryptType"`
	MerchantID        string `json:"MerchantID" db:"MerchantID"`
	CheckMacValue     string `json:"CheckMacValue" db:"CheckMacValue"`
	HashKey           string `json:"HashKey" db:"HashKey"`
	HashIV            string `json:"HashIV" db:"HashIV"`
}

func (c *Cashflow) FindByOid(order_id int) (err error) {
	db, err := NewDBSolomo()

	if err != nil {
		log.Printf("connection error")
	}
	err = db.Where("order_id = ?", order_id).First(c)
	if err != nil {
		log.Printf(`order_id: %v is not exist`, err)
		return
	}
	return
}

func (e *EcpayCashFlow) GetEcpayParams(orderNo string) (err error) {

	o := new(Order)
	// orderNo exist in DB
	findNoting := o.FindByNo(orderNo)
	defaultMemo := "new orderNo"
	noMemo := "no memo"

	if findNoting != nil {
		o.CreatedAt = time.Now()
		o.Total = 999
		o.Remark = &defaultMemo
	}
	if o.Remark == nil {
		o.Remark = &noMemo
	}
	// get ItemNames' array
	pName := GetItemNames(o.ID)

	// 測試店金額上限為30000
	if o.Total >= 30000 {
		o.Total = 29999
	}

	// ecpay params
	HashKey := "HashKey=5294y06JbISpM5x9"
	HashIV := "HashIV=v77hoKGq4kWxNNIS"
	e.ChoosePayment = "ALL"
	e.EncryptType = "1"
	e.ItemName = strings.Join(pName, ",")
	e.MerchantID = "2000132"
	e.MerchantTradeDate = o.CreatedAt.Format("2006/01/02 15:04:05")
	e.MerchantTradeNo = orderNo
	e.PaymentType = "aio"
	e.ReturnURL = "https://upayment.xinmedia.com/api/ecpay_res"
	e.TotalAmount = strconv.FormatFloat(o.Total, 'f', 0, 64)
	e.TradeDesc = *(o.Remark)

	arr := []string{"ChoosePayment", "EncryptType", "ItemName", "MerchantID", "MerchantTradeDate", "MerchantTradeNo", "PaymentType", "ReturnURL", "TotalAmount", "TradeDesc"}
	t := reflect.ValueOf(*e)

	str := HashKey + "&"

	for _, item := range arr {
		str += item + "=" + t.FieldByName(item).Interface().(string) + "&"
	}

	str += HashIV

	// 轉為.Net encode結果
	encodeStr := url.QueryEscape(str)
	encodeStr = strings.ToLower(encodeStr)
	encodeStr = strings.ReplaceAll(encodeStr, "%21", "!")
	encodeStr = strings.ReplaceAll(encodeStr, "-", "%7e")
	encodeStr = strings.ReplaceAll(encodeStr, "%2A", "*")
	encodeStr = strings.ReplaceAll(encodeStr, "%28", "(")
	encodeStr = strings.ReplaceAll(encodeStr, "%29", ")")

	// sha256 加密
	hash := sha256.New()
	hash.Write([]byte(encodeStr))
	md := hash.Sum(nil)
	macValue := hex.EncodeToString(md)

	macValue = strings.ToUpper(macValue)

	e.CheckMacValue = macValue

	return nil
}

// String is not required by pop and may be deleted
func (c Cashflow) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Cashflows is not required by pop and may be deleted
type Cashflows []*Cashflow

// String is not required by pop and may be deleted
func (c Cashflows) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Cashflow) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Cashflow) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Cashflow) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
