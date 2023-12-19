package models

import (
	"blog_server/global"
	"blog_server/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 33, 11, 51, 111, 200, 255, 10}
	key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		global.LOG.Warningf("密码加密错误: %v", err)
	}
	pw := base64.StdEncoding.EncodeToString(key)
	return pw
}

// VerifyLogin 登录验证
func VerifyLogin(username string, password string) int {
	var user User
	global.DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_PERMISSION
	}
	return errmsg.SUCCESS
}

// BeforeSave 钩子函数
func (user *User) BeforeSave(db *gorm.DB) error {
	user.Password = ScryptPw(user.Password)
	return nil
}

// CheckUser 根据username查询用户是否存在
func CheckUser(name string) int {
	var user User
	global.DB.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

func CheckUserFromId(id int) int {
	var user User
	affected := global.DB.First(&user, id).RowsAffected
	if affected == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// CreateUser 添加用户
func CreateUser(user *User) int {
	//user.Password = ScryptPw(user.Password)
	err := global.DB.Create(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 获取用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	global.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	return users, errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	if global.DB.Where("id =?", id).First(&user).RowsAffected < 1 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	err := global.DB.Where("id = ?", id).Delete(&user).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUserInfo 获取用户详细信息
func GetUserInfo(id int) (User, int) {
	var user User
	err := global.DB.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCESS
}

// UpdateUser 更新用户
func UpdateUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := global.DB.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
