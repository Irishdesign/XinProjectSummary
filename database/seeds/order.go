package main

import (
	"XinAPI/app/http/models"
	"XinAPI/app/http/services"
	models_gen "XinAPI/build/gqlgen/admin/models"
	"XinAPI/pkg/l"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gobuffalo/pop"
)

func main() {
	vendorDB := getVendorData()

	db, err := models.NewDBSolomo()

	_platform := []string{"XINSHOP", "XINTUKU"}

	_firstName := []string{"小明", "小華", "小珠"}
	_lastName := []string{"王", "陳", "林"}
	for i := 1; i <= 100; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		c := "XIN"
		data := models.Order{
			Status:        rand.Intn(3),
			Platform:      _platform[rand.Intn(len(_platform))],
			Type:          rand.Intn(2),
			MemberID:      rand.Intn(100000),
			FirstName:     _firstName[rand.Intn(len(_firstName))],
			LastName:      _lastName[rand.Intn(len(_lastName))],
			Email:         "test@gmail.com",
			Country:       "TW",
			AddressCity:   "TPE",
			AddressArea:   "TPE",
			CouponID:      &c,
			Address:       "address 168",
			CountryCode:   "886",
			Mobile:        "954658876",
			Total:         0,
			PaymentType:   rand.Intn(100),
			PaymentStatus: rand.Intn(1000),
			InvoiceType:   rand.Intn(2),
		}

		err = db.Save(&data)
		l.Log(`err----- addOrder`, err)
		l.Log(`訂單ID=`, data.ID)
		totalprice := addOrderVendor(data.ID, vendorDB)

		query := db.Where("id = ?", data.ID)
		data.Total = totalprice
		query.Connection.UpdateColumns(&data, "total")
		addStatus(data.ID)
	}

}

func addStatus(orderid int) {
	db, err := models.NewDBSolomo()
	data := models.OrderStatus{
		Status:  rand.Intn(2),
		OrderID: orderid,
	}
	// l.Log(`訂單ITEM資料`, data)
	err = db.Save(&data)
	l.Log(`err-----`, err)
}

func addItem(orderid int, vendorid int) float64 {
	// ProductDB := getProductData()
	db, err := models.NewDBSolomo()

	total := 0.0
	productArr := FindProductByVid(db, vendorid)
	totalProd := len(*productArr)

	if totalProd == 0 {
		addProduct(vendorid)
		productArr = FindProductByVid(db, vendorid)
		totalProd = len(*productArr)
	}

	for i := 1; i <= 2; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		_productFrom := []string{"Video", "Photo", "Travel"}
		_id := rand.Intn(totalProd)
		data := models.OrderItem{
			OrderVendorID:  vendorid,
			OrderID:        orderid,
			ProductFrom:    _productFrom[rand.Intn(len(_productFrom))],
			ProductID:      (*productArr)[_id].ID,
			Qty:            rand.Intn(4) + 1,
			Price:          *(*productArr)[_id].SalePrice,
			CouponID:       orderid,
			ParentID:       0,
			ShipmentStatus: 1,
		}
		// l.Log(`訂單ITEM資料`, data)
		err = db.Save(&data)
		l.Log(`err-----addItem`, err)
		total += data.Price * float64(data.Qty)
	}
	return total
}

func FindProductByVid(db *pop.Connection, vendorId int) *models.Products {
	m := new(models.Products)
	db.Where(`vendor_id = ?`, vendorId).All(m)
	return m
}

func addOrderVendor(orderid int, vendorDB *models.Vendors) float64 {
	db, err := models.NewDBSolomo()
	total := 0.0
	basicNum := orderid % 10

	for i := 1; i <= 3; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		data := models.OrderVendor{
			OrderID:  orderid,
			VendorID: (*vendorDB)[(basicNum+i)%10].ID,
		}

		err = db.Save(&data)
		l.Log(`err----- addOrderVendor`, err)
		subtotal := addItem(orderid, data.VendorID)
		total += subtotal
	}
	return total
}

func getVendorData() *models.Vendors {
	db, _ := models.NewDBSolomo()
	vendorDB := new(models.Vendors)
	db.All(vendorDB)

	if len(*vendorDB) == 0 {
		addVendor()
		db.All(vendorDB)
		return vendorDB
	}
	return vendorDB

}
func getProductData() *models.Products {
	db, _ := models.NewDBSolomo()
	productDB := new(models.Products)
	db.All(productDB)
	if len(*productDB) == 0 {
		addProduct(0)
		db.All(productDB)
		return productDB
	}
	return productDB
}

func addVendor() {

	db, err := models.NewDBSolomo()
	_vendor := []int{}
	for i := 1; i <= 10; i++ {
		_vendor = append(_vendor, rand.Intn(9999999))
	}
	data := models.Vendor{}
	for _, vendor := range _vendor {
		data = models.Vendor{
			No:   vendor,
			Name: getCname(vendor, 0),
			Bin:  getCname(vendor, 1),
		}
		err = db.Save(&data)
		l.Log(`err- addVendor----`, err)
	}

}

func getCname(id int, ty int) *string {
	s := strconv.Itoa(id)
	arr := []string{"霆", "曉", "衡", "儒", "靜", "翰", "蔚", "憶", "雙", "濤", "麗", "韻", "耀", "藝", "巍", "蘭", "雪", "堯", "誼", "影", "慧", "潔", "潤", "成", "翔", "隆", "東", "森", "迪", "賽", "睿", "艾", "高", "德", "雅", "格", "納", "欣", "億", "維", "銳", "菲", "佳", "沃", "晟", "捷", "樂", "飛", "福", "皇", "嘉", "達", "佰", "美", "元", "亮", "名", "歐", "特", "辰", "康", "訊", "鵬", "騰", "宏", "偉", "鈞", "思", "正"}
	title := ""
	if ty == 0 {
		for i := 1; i < 3; i++ {
			r := rand.Intn(len(arr))
			title += arr[r]
		}
		title += "股份有限公司"
		return &title
	}
	return &s

}
func addProduct(vendorid int) {
	vendorDB := getVendorData()
	rand.Seed(time.Now().UTC().UnixNano())
	run := 30
	if vendorid != 0 {
		run = 3
	}
	for i := 1; i <= run; i++ {
		Name := fmt.Sprintf("Order商品(%03d)", i)
		Brief := fmt.Sprintf("我是商品(%03d)", i)
		Desp := fmt.Sprintf("<p>我是商品(%03d)</p>", i)
		layout := "2006-01-02 15:04:05"
		input := fmt.Sprintf("2020-01-01 01:%02d:%02d", i/10, i%10)

		PublishSdate, Terr := time.Parse(layout, input)
		if Terr != nil {
			return
		}

		Price := (rand.Intn(23) + 1) * 500
		ListPrice := float64(Price)
		SalePrice := float64(Price) * 0.8
		vendorID := rand.Intn(len(*vendorDB))
		if vendorid != 0 {
			vendorID = vendorid
		}

		Data := models.Product{
			Name:         Name,
			Type:         1,
			State:        0,
			VendorID:     vendorID,
			ListPrice:    &ListPrice,
			SalePrice:    &SalePrice,
			Brief:        &Brief,
			Desp:         &Desp,
			PublishSdate: PublishSdate,
		}

		// Create
		Data.Save()

		// Update no

		Data.CreateNo("P-%07d")

		var Size1, Size2, Color1, Color2 *models_gen.IProductSpec
		var Size, Spec []*models_gen.IProductSpec
		ONE := 1
		Size1 = &models_gen.IProductSpec{
			Name:      "Big",
			SalePrice: &SalePrice,
			Qty:       &ONE,
			Vw:        &ONE,
		}

		Size2 = &models_gen.IProductSpec{
			Name:      "Small",
			SalePrice: &SalePrice,
			Qty:       &ONE,
			Vw:        &ONE,
		}

		if i%3 == 2 {
			Size = []*models_gen.IProductSpec{Size1, Size2}
		}

		Color1 = &models_gen.IProductSpec{
			Name:      "RED",
			SalePrice: &SalePrice,
			Qty:       &ONE,
			Vw:        &ONE,
			Items:     Size,
		}

		Color2 = &models_gen.IProductSpec{
			Name:      "BLUE",
			SalePrice: &SalePrice,
			Qty:       &ONE,
			Vw:        &ONE,
			Items:     Size,
		}
		if i%3 != 0 {
			Spec = []*models_gen.IProductSpec{Color1, Color2}
		}

		services.AddSpec(Spec, Data.ID, 0)

		fmt.Printf("Add ID=%03d \n", Data.ID)
	}

}
