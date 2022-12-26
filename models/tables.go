package models

var Tables = []interface{}{
	&User{},
	&Account{},
}

type User struct {
	Id        string    `json:"id" gorm:"primarykey"`
	UserName  string    `json:"user_name" gorm:"type:varchar(100)"`
	FirstName string    `json:"first_name" gorm:"type:varchar(100)"`
	LastName  string    `json:"last_name" gorm:"type:varchar(100)"`
	Accounts  []Account `json:"accounts" gorm:"foreignkey:UserId;references:Id"`
}

type Account struct {
	Id            string  `json:"id" gorm:"primarykey"`
	UserId        string  `json:"user_id" gorm:"type:varchar(100)"`
	AccountNumber string  `json:"account_number" gorm:"type:varchar(100)"`
	Balance       float64 `json:"balance" gorm:"type:float"`
}
