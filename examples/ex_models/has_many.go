package ex_models

// User has many CreditCards, UserID is the foreign key
type User struct {
	ID          uint         `json:"id" gorm:"primary_key"`
	Name        string       `json:"name"`
	CreditCards []CreditCard `json:"credit_cards"`
}

type CreditCard struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Number string `json:"number"`
	UserID uint   `json:"user_id"`
}
