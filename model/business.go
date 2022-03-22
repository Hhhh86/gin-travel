package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Business 旅游公司
type Business struct {
	gorm.Model
	CompanyName    string
	PasswordDigest string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

// SpotInfo 景点信息
type SpotInfo struct {
	gorm.Model
	Name         string
	Address      string
	AgeScope     string //0Children、1Youth、2Midlife，3Aged
	Introduction string
	Tips         string
	OpenTime     string
	PlaceType    int //1自然生态类、2历史文化类、3现代游乐类、4产业融合类、5其他类。
	Price        int //门票价格
}

// GetUser 用ID获取用户
func GetBusiness(ID interface{}) (Business, error) {
	var business Business
	result := DB.First(&business, ID)
	return business, result.Error
}

// SetPassword 设置密码
func (user *Business) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *Business) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
