package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Firts_name      *string            `json:"firts_name"`
	Last_name       *string            `json:"last_name"`
	Password        *string            `json:"password"`
	Email           *string            `json:"email"`
	Phone           *string            `json:"phone"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	User_id         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"user_cart" bson:"usercart"`
	Address_Details []Address          `json:"address_details" bson:"address"`
	Order_Status    []Order            `json:"order_status" bson:"orders"`
}
type Product struct {
	Product_Id   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        uint64             `json:"price"`
	Rating       *uint64            `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_Id   primitive.ObjectID ` bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_Cart     []ProductUser
	Ordered_At     *time.Time
	Price          int
	Discount       *int
	Payment_Method Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
