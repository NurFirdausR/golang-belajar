package belajargolanggorm

import "time"

// User => users
// OrderDetail => order_details

type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:created_at;autoCreateTime;autoUpdatedTime"`
	Information string    `gorm:"-"` // - artinya tidak ada di databse
	Wallet      Wallet    `gorm:"foreignKey:user_id;references:id"`
}

type Name struct { //embedded
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

func (u *User) TableName() string {
	return "users"
}
