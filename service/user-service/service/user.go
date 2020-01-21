package service

import (
	"github.com/jinzhu/gorm"
	"pea-web/service/user-service/model"
	"pea-web/api/plus"
	"pea-web/service/user-service/repositories"
	"pea-web/api/tools"
	"strings"
	"time"
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
		return nil, plus.PE_NotFoundUser
	}
	if err := tools.IsValidateUsername(username); err != nil {
		return nil, plus.PE_UserFormatError
	}
	if err := tools.IsValidatePassword(password, repassword); err != nil {
		return nil, plus.PE_PassWordError
	}
	if len(email) > 0 {
		if err := tools.IsValidateEmail(email); err != nil {
			return nil, plus.PE_EmailFormatError
		}
		if s.isEmailExists(email) != nil {
			return nil, plus.PE_EmailHasOccupy
		}
	}
	if s.isUserNameExists(username) != nil {
		return nil, plus.PE_UserNameOccupy
	}

	user := &model.User{
		Username:   username,
		Email:      email,
		Nickname:   nickname,
		Password:   password,
		Status:     0,
		CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6,
	}
	err := tools.Tx(cmd.DB, func(tx *gorm.DB) error {
		//插入用户
		if err := repositories.UserRepository.Create(tx, user); err != nil {
			return err
		}
		/*
			//获取头像
			avatarUrl, err := s.GetAvatar(user.ID)
			if err != nil {
				return err
			}
			// 更新头像
			if err := repositories.UserRepository.UpdataUpAvatar(tx, user.ID, avatarUrl); err != nil {
				return err
			}
		*/
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 查询邮箱
func (s *userService) isEmailExists(email string) *model.User {
	return repositories.UserRepository.GetByEmail(cmd.DB, email)
}

//查询用户名
func (s *userService) isUserNameExists(username string) *model.User {
	return repositories.UserRepository.GetByUserName(cmd.DB, username)
}

func (s *userService) CheckUsernameExists(username string) bool {
	return repositories.UserRepository.GetByUserName(cmd.DB, username) != nil
}

//第三方查询
func (s *userService) Get(id int64) *model.User {
	return repositories.UserRepository.Get(cmd.DB, id)
}

//更新用户表
func (s *userService) Update(id int64, name string, value interface{}) error {
	return repositories.UserRepository.Updata(cmd.DB, id, name, value)
}

//获取头像
func (s *userService) GetAvatar(id int) (string, error) {
	avatarBytes, err := tools.Generate(int64(id))
	if err != nil {
		return "", err
	}
	return tools.UploadImage(avatarBytes)
}

// ************   逻辑 ***************************

//登录
func (s *userService) Login(username, password string) (*model.User, error) {
	if len(username) == 0 {
		return nil, plus.PE_IsNotEmptyName
	}
	if len(password) == 0 {
		return nil, plus.PE_PassWordError
	}

	var user *model.User = nil
	if err := tools.IsValidateEmail(username); err != nil {
		user = s.isEmailExists(username)
	} else {
		user = s.isUserNameExists(username)
	}
	if user == nil {
		return nil, plus.PE_UserNameOccupy
	}
	//判断密码
	if password != user.Password {
		return nil, plus.PE_PassWordError
	}
	return user, nil
}

//第三方账号登录
func (s *userService) LoginByThirdAccount(thridAccount *model.ThirdAccount) (*model.User, error) {
	user := s.Get(int64(thridAccount.UserId))
	if user != nil {
		return user, nil
	}
	user = &model.User{
		Nickname:   thridAccount.Nickname,
		Status:     0,
		CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6,
	}
	err := tools.Tx(cmd.DB, func(tx *gorm.DB) error {
		if err := repositories.UserRepository.Create(tx, user); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, plus.PE_ThirdAccountError
	}
	return user, nil
}

//设置用户名
func (s *userService) SetUserName(user_id int64, user_name string) error {
	user_name = strings.TrimSpace(user_name)
	if err := tools.IsValidateUsername(user_name); err != nil {
		return plus.PE_UserFormatError
	}
	user := s.Get(user_id)
	if len(user.Username) > 0 {
		return plus.PE_UserNameHasSet
	}
	if s.CheckUsernameExists(user_name) {
		return plus.PE_UserNameOccupy
	}
	if err := s.Update(user_id, "username", user_name); err != nil {
		return plus.PE_UpdateUserNameError
	}
	return nil
}
