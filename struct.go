package main

import "fmt"

type Address struct {
	AddressID     string `json:"addressId" bson:"addressId"`
	AddFirstLine  string `json:"addFirstLine" bson:"addFirstLine"`
	AddSecondLine string `json:"addSecondLine" bson:"addSecondLine"`
	City          string `json:"city" bson:"city"`
	State         string `json:"state" bson:"state"`
	Country       string `json:"country" bson:"country"`
	ZipCode       string `json:"zipCode" bson:"zipCode"`
	Phone         string `json:"phone" bson:"phone"`
	Fax           string `json:"fax" bson:"fax"`
	AddressType   string `json:"addressType" bson:"addressType"`
}

type PurchaseOrder struct {
	PurchaseOrderID   string                         `bson:"poId"`
	Amount            float64                        `bson:"amount"`
}

type Account struct {
	name string `bson:"name"`
	newName *string `bson:"newName"`
	CompanyAddress   *Address         `bson:"companyAddress"`
	PurchaseOrders   PurchaseOrder        `bson:"purchaseOrders"`
}

func main()  {
	newAcc := Account{}
	fmt.Println("Account name: ", newAcc.name)
	fmt.Println("Account newName: ", newAcc.newName)
	fmt.Println("Account company address: ", newAcc.CompanyAddress)
	//fmt.Println("Account company address addressId: ", newAcc.CompanyAddress.AddressID)
	//fmt.Println("Account company address city: ", &newAcc.CompanyAddress.City)
	fmt.Println("Account purchase order: ", newAcc.PurchaseOrders)
	fmt.Println("Account purchase order amount: ", newAcc.PurchaseOrders.Amount)
	fmt.Println("Account purchase order purchaseOrderId: ", newAcc.PurchaseOrders.PurchaseOrderID)
}
