package ex_models

// User has one CreditCardOne, CreditCardID is the foreign key

type UserOne struct {
	CreditCard CreditCard
}

type CreditCardOne struct {
	Number string
	UserID uint
}
