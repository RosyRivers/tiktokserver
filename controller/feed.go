package controller

import (
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var VideoList []Video
	videos, err := repository.NewVideoDaoInstance().QueryVideoFeed()
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Fail to query videos!\n"},
		})
	}
	VideoList, err = ConvertVideoDBToJSON(videos)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "Videos are found, but Convert is failed!\n"},
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: VideoList,
		NextTime:  time.Now().Unix(),
	})
}
