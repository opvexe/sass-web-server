package repositories

import (
	"github.com/jinzhu/gorm"
	"pea-web/api/model"
)

type userRepository struct {
}

func newRepository() *userRepository {
	return new(userRepository)
}

//初始化
var UserRepository = newRepository()


// 查询邮箱
func (r *userRepository)GetByEmail(db *gorm.DB,email string) *model.User{
	var user  model.User
	if err:=db.Where("email=?",email).Find(&user).Error;err!=nil{
		return nil
	}
	return &user
}

// 查询用户
func (r *userRepository) GetByUserName(db *gorm.DB,username string) *model.User{
	var user model.User
	if err:=db.Where("username=?",username).Find(&user).Error;err!=nil{
		return nil
	}
	return &user
}

//注册
func (r *userRepository) Create(db *gorm.DB,user *model.User) error {
	return db.Create(user).Error
}

//更新用户头像
func (r *userRepository) UpdataUpAvatar(db *gorm.DB,id int,avatar string) error {
	return db.Model(new(model.User)).Where("id=?",id).Update("avatar",avatar).Error
}