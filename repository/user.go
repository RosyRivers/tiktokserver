package repository

import (
	"errors"
	// "github.com/RaymondCode/simple-demo/util"
	"fmt"
	"gorm.io/gorm"
	"sync"
	// "time"
)

type User struct {
	Id            int64  `gorm:"column:id" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	Token         string `gorm:"column:token" json:"token"`
	FollowCount   int64  `gorm:"column:followcount" json:"follow_count"`
	FollowerCount int64  `gorm:"column:followercount" json:"follower_count"`
	IsFollow      bool   `gorm:"column:isfollow" json:"is_follow"`
}

func (User) TableName() string {
	return "users"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserByToken(token string) (*User, error) {
	var user User
	err := db.Where("token = ?", token).First(&user).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, err
	}
	if err != nil {
		fmt.Println(err.Error())
		// util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		// util.Logger.Error("insert post err:" + err.Error())
		return err
	}
	return nil
}
