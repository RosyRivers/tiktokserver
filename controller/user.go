package controller

import (
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]repository.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User repository.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	_, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		isfollow := false
		newUser := &repository.User{
			Id:            userIdSequence,
			Name:          username,
			Token:         token,
			FollowCount:   1,
			FollowerCount: 0,
			IsFollow:      isfollow,
		}
		if err := repository.NewUserDaoInstance().CreateUser(newUser); err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "insert user err:" + err.Error()},
			})
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Find user by token error"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	user, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
