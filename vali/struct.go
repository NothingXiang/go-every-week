package vali

type TstStruct1 struct {
	// 强制相等
	Const string `validate:"eq=const"`
}

type Address struct {
}

type User struct {
	FirstName   string     `validate:"gt=0"`
	LastName    string     `validate:"required"`
	Age         uint8      `validate:"gte=0,lte=80"`
	Email       string     `validate:"required,email"`
	FavourColor string     `validate:"iscolor"`
	Address     []*Address `validate:"required,dive,required"`
}
