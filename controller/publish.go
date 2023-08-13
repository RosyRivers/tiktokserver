package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
	Token     string  `json:"token"`
}

var videoIdSequence = int64(0)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	fmt.Println(token)
	// date  := c.PostForm("date")

	user, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(title)
	id, err := repository.NewVideoDaoInstance().QueryVideoLatest()
	if err != nil {
		return
	}
	if id == -1 {
		fmt.Println("query error:", err)
		return
	}
	videoIdSequence = id
	atomic.AddInt64(&videoIdSequence, 1)
	// user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d-%d-%s.mp4", user.Id, videoIdSequence, filename)
	saveFile := filepath.Join("./public/video", finalName)

	videourl := []string{"http://192.168.1.9:8080/static/video", finalName}
	playurl := strings.Join(videourl, "/")
	CoverName := fmt.Sprintf("%spng", finalName[0:len(finalName)-3])
	pngurl := []string{"http://192.168.1.9:8080/static/cover", CoverName}
	coverurl := strings.Join(pngurl, "/")
	newVideo := &repository.Video{
		Token:         token,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		UploadTime:    time.Now(),
	}
	if err := repository.NewVideoDaoInstance().CreateVideo(newVideo); err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "insert video err:" + err.Error()},
		})
	}
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	utils.SaveCover(saveFile, finalName)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	_, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "User not exist\n"},
		})
	}
	videos, err := repository.NewVideoDaoInstance().QueryVideoByAuthor(token)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "query video list err:" + err.Error(),
			},
		})
	}
	VideoList, err := ConvertVideoDBToJSON(videos)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "videos are found, but Convert is failed"},
		})
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},

		VideoList: VideoList,
		Token:     token,
	})
}
