package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"pea-web/api/model"
	"pea-web/api/repositories"
	"pea-web/api/tools"
	"pea-web/cmd"
)

type userService struct {
}

func newUserService() *userService {
	return new(userService)
}

//初始化
var UserService = newUserService()

//注册
func (s *userService) Register(username, email, nickname, password, repassword string) (*model.User, error) {
	if len(nickname) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if err := tools.IsValidateUsername(username); err != nil {
		return nil, errors.New("用户名格式无效")
	}
	if err := tools.IsValidatePassword(password, repassword); err != nil {
		return nil, errors.New("用户密码输入错误")
	}
	if len(email) > 0 {
		if err := tools.IsValidateEmail(email); err != nil {
			return nil, errors.New("邮箱格式不正确")
		}
		if s.isEmailExists(email) {
			return nil, errors.New("邮箱被占有")
		}
	}
	if s.isUserNameExists(username) {
		return nil, errors.New("用户名被占有")
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Nickname: nickname,
		Password: password,
		Status:   0, //0
	}
	err := tools.Tx(cmd.DB, func(tx *gorm.DB) error {
		//插入用户
		if err := repositories.UserRepository.Create(tx, user); err != nil {
			return err
		}
		//获取头像
		avatarUrl, err := s.GetAvatar(user.ID)
		if err != nil {
			return err
		}
		// 更新头像
		if err := repositories.UserRepository.UpdataUpAvatar(tx, user.ID, avatarUrl); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 查询邮箱
func (s *userService) isEmailExists(email string) bool {
	return repositories.UserRepository.GetByEmail(cmd.DB, email) != nil
}

//查询用户名
func (s *userService) isUserNameExists(username string) bool {
	return repositories.UserRepository.GetByUserName(cmd.DB, username) != nil
}

//获取头像
func (s *userService) GetAvatar(id int) (string, error) {
	avatarBytes, err := tools.Generate(int64(id))
	if err != nil {
		return "", err
	}
	return tools.UploadImage(avatarBytes)
}

//登录
func (s *userService) Login(username ,password string) (*model.User, error) {
	if len(username) == 0 {
		return nil,errors.New("用户名/邮箱不能为空")
	}
	if len(password) == 0 {
		return nil,errors.New("密码不能为空")
	}
	if s.isUserNameExists(username) {
		return nil, errors.New("用户名被占有")
	}
}

//获取token
func (s *userService) Generate (id int)  {

}