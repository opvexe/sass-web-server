package service

import (
	"github.com/jinzhu/gorm"
	"pea-web/api/model"
	"pea-web/api/repositories"
	"pea-web/cmd"
	"pea-web/tools"
)

type userService struct {
}

func newUserService() *userService {
	return new(userService)
}

//初始化
var UserService = newUserService()

//注册
func (s *userService) Register(username, email, nickname, password, repassword string) (*model.User, string) {
	if len(nickname) == 0 {
		return nil, tools.RECODE_NODATA
	}
	if err := tools.IsValidateUsername(username); err != nil {
		return nil, tools.RECODE_PARAMERR
	}
	if err:=tools.IsValidatePassword(password,repassword);err!=nil {
		return nil, tools.RECODE_PARAMERR
	}
	if len(email)>0 {
		if err:=tools.IsValidateEmail(email);err!=nil{
			return nil, tools.RECODE_PARAMERR
		}
		if s.isEmailExists(email) {	//邮箱被占用
			return nil, tools.RECODE_PARAMERR
		}
	}
	if s.isUserNameExists(username){	//用户名被占用
		return nil, tools.RECODE_PARAMERR
	}

	user := &model.User{
		Username:    username,
		Email:       email,
		Nickname:    nickname,
		Password:    password,
		Status:      0, //0
	}
	err := tools.Tx(cmd.DB, func(tx *gorm.DB) error {
		//插入用户
		if err := repositories.UserRepository.Create(tx,user);err!=nil{
			return err
		}
		//获取头像
		avatarUrl,err := s.GetAvatar(user.ID)
		if err!=nil{
			return err
		}
		// 更新头像
		if err :=repositories.UserRepository.UpdataUpAvatar(tx,user.ID,avatarUrl);err!=nil{
			return err
		}
		return nil
	})
	if err!=nil {
		return nil, tools.RECODE_DBERR
	}
	return user, tools.RECODE_OK
}










// 查询邮箱
func (s *userService)isEmailExists(email string) bool {
	return repositories.UserRepository.GetByEmail(cmd.DB,email)!=nil
}


//查询用户名
func (s *userService)isUserNameExists(username string) bool {
	return repositories.UserRepository.GetByUserName(cmd.DB,username)!=nil
}

//获取头像
func (s *userService) GetAvatar(id int) (string,error){
	avatarBytes,err := tools.Generate(int64(id))
	if err!=nil {
		return "",err
	}
	return tools.UploadImage(avatarBytes)
}