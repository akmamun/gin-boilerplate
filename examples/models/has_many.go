package models

// User has many CreditCards, UserID is the foreign key
type User struct {
	CreditCards []CreditCard `json:"credit_cards"`
}

type CreditCard struct {
	Number string
	UserID uint
}
