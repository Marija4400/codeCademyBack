package models

type User struct {
	Common
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   uint   `json:"roleId"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleId"`
}
