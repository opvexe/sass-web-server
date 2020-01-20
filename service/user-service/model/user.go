package model

//用户表
type User struct {
	ID          int
	Username    string `gorm:"size:32;unique;" json:"username" form:"username"`
	Email       string `gorm:"size:128;unique;" json:"email" form:"email"`
	Nickname    string `gorm:"size:16;" json:"nickname" form:"nickname"`
	Avatar      string `gorm:"type:text" json:"avatar" form:"avatar"`
	Password    string `gorm:"size:512" json:"password" form:"password"`
	Status      int    `gorm:"index:idx_status;not null" json:"status" form:"status"`
	Roles       string `gorm:"type:text" json:"roles" form:"roles"`
	Type        int    `gorm:"not null" json:"type" form:"type"`
	Description string `gorm:"type:text" json:"description" form:"description"`
	CreateTime  int64  `json:"createTime" form:"createTime"`
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`
}

//用户token
type UserToken struct {
	ID         int
	Token      string `gorm:"size:32;unique;not null" json:"token" form:"token"`
	UserId     int64  `gorm:"not null;index:idx_user_id;" json:"userId" form:"userId"`
	ExpiredAt  int64  `gorm:"not null" json:"expiredAt" form:"expiredAt"`
	Status     int    `gorm:"not null;index:idx_status" json:"status" form:"status"`
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`
}
