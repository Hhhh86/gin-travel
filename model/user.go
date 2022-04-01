package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"index:idx_name,unique"`
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

// UserInfo 用户详细信息模型
type UserInfo struct {
	gorm.Model
	Sex            int //1女 2男
	PlaceType      int //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	AgeScope       int //0Children、1Youth、2Midlife，3Aged
	ExpectMinPrice int //期望最小价格
	ExpectMaxPrice int //期望最大价格
	UserName       string
	NickName       string
	Avatar         string `gorm:"size:1000"`
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
