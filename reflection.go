package main

import (
	"fmt"
	"reflect"
)
type UpdateCompanyAddressDto struct {
	AddFirstLine  string `json:"addFirstLine" bson:"addFirstLine"`
	AddSecondLine string `json:"addSecondLine" bson:"addSecondLine"`
	City          string `json:"city" bson:"city"`
	State         string `json:"state" bson:"state"`
	Country       string `json:"country" bson:"country"`
	ZipCode       string `json:"zipCode" bson:"zipCode"`
	Phone         string `json:"phone" bson:"phone"`
	Fax           string `json:"fax" bson:"fax"`
}
type PatchArray struct {
	Replace interface{} `json:"replace"` // if replace given all others are ignored
	Add     interface{} `json:"add"`
	Remove  interface{} `json:"remove"`
	Update  interface{} `json:"update"`
}
type UpdateAccountDTO struct {
	CompanyAddress *UpdateCompanyAddressDto `json:"companyAddress"`
	PurchaseOrders *PatchArray         `json:"purchaseOrders"`
}

type UpdatePurchaseOrder struct {
	PurchaseOrderID string           `json:"poId"`
	InvoiceID       string           `json:"invoiceId"`
	OrderStatus     string   `json:"status"`
}

func main(){
	updatePO := &UpdateAccountDTO{
		PurchaseOrders: &PatchArray{
			Update: []UpdatePurchaseOrder{{
				OrderStatus: "Received",
			}},
		},
	}

	ptrVal := reflect.ValueOf(updatePO).Elem();
	fmt.Printf("Value of ptrVal: %s", ptrVal)

}
