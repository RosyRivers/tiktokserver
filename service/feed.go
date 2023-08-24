package service

import (
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	//"github.com/u2takey/go-utils/json"

	"encoding/json"
	// "strconv"
	"time"
)

type FeedResponse struct {
	repository.Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(username string) (FeedResponse, error) {

	var VideoList []Video
	err := repository.NewVideoDaoInstance().UpdateVideoUrl(serverDomain)
	if err != nil {
		feedResponse := FeedResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  "fail to update url!\n"},
		}
		return feedResponse, err
	}
	videos, err := repository.NewVideoDaoInstance().QueryVideoFeed()
	if err != nil {
		feedResponse := FeedResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  "Fail to query videos!\n"},
		}
		return feedResponse, err
	}

	//已登陆用户更新视频的点赞状态
	if username != "" {
		likedVideosId, err := repository.NewFavouriteInstance().FindUserLikedVideo(username)
		if err != nil {
			_ = utils.WriteLog("feed_querylike.txt", err.Error())
			feedResponse := FeedResponse{
				Response: repository.Response{
					StatusCode: 1,
					StatusMsg:  "Fail to query liked videos!\n" + err.Error()},
			}
			return feedResponse, err
		}

		likedVideosId_map := utils.Tomap(likedVideosId)

		for _, video := range *videos {
			_, isliked := likedVideosId_map[int(video.Id)]
			if isliked == true {
				video.IsFavorite = true
			}
		}

	}
	jsondata, _ := json.Marshal(videos)
	_ = utils.WriteLog("is_favorite.txt", string(jsondata))
	VideoList, err = ConvertVideoDBToJSON(videos)

	if err != nil {
		feedResponse := FeedResponse{
			Response: repository.Response{StatusCode: 1,
				StatusMsg: "Videos are found, but Convert is failed!\n"},
		}
		return feedResponse, err
	}
	feedResponse := FeedResponse{
		Response:  repository.Response{StatusCode: 0},
		VideoList: VideoList,
		NextTime:  time.Now().Unix(),
	}
	return feedResponse, nil
}