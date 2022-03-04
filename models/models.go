package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firts_name      *string
	Last_name       *string
	Password        *string
	Email           *string
	Phone           *string
	Token           *string
	Refresh_Token   *string
	Created_at      time.Time
	Updated_at      time.Time
	User_id         string
	UserCart        []ProductUser
	Address_Details []Address
	Order_Status    []Order
}
type Product struct {
	Product_Id
	Product_Name
	Price
	Rating
	Image
}

type ProductUser struct {
	Product_Id
	Product_Name
	Price
	Rating
	Image
}

type Address struct {
	Address_id
	House
	Street
	City
	Pincode
}

type Order struct {
	Order_ID
	Order_Cart
	Ordered_At
	Price
	Discount
	Payment_Method
}

type Payment struct {
	Digital bool
	COD     bool
}
