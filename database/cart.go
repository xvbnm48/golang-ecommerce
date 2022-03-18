package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdNotValid     = errors.New(" user id not valid")
	ErrCantUpdateUser     = errors.New(" can't add the product to cart")
	ErrCantRemoveItemCart = errors.New("cant remove item from cart")
	ErrCantGetItem        = errors.New("was unable to get item from the cart")
	ErrCantBuyItem        = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuy() {

}
