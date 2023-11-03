// model/user.go
package model

const UserTable = "users"

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`           // 用户ID
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"` // 用户邮箱(唯一)
	Name     string `json:"name" gorm:"type:varchar(255);not null"`         // 用户名
	Password string `json:"password" gorm:"type:varchar(255);not null"`     // 用户密码

	// edge
	Links []Link `json:"-" gorm:"foreignKey:OwnerID;references:ID"`
}
