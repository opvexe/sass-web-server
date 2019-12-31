package model

var Model = []interface{}{
	new(User),
}

// 用户表
type User struct {
	ID          int
	Username    string `gorm:"size:32;unique"`
	Email       string `gorm:"size:128;unique"`
	Nickname    string `gorm:"size:16"`
	Avatar      string `gorm:"size:256"`
	Password    string `gorm:"size:512"`
	Status      int    `gorm:"index:idx_status;not null"`
	Roles       string `gorm:"type:text"`
	Type        int    `gorm:"not null"`
	Description string `gorm:"type:text"`
}
