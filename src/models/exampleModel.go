package models

type Example struct {
	Data string `json:"data"`
	//CreatedAt *time.Time `gorm:"type:timestamp" json:"created_at,string,omitempty"`
	//UpdatedAt *time.Time `gorm:"type:timestamp" json:"updated_at_at,string,omitempty"`
}
