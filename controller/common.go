package controller

import (
	"github.com/RaymondCode/simple-demo/repository"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64           `json:"id,omitempty"`
	Author        repository.User `json:"author"`
	PlayUrl       string          `json:"play_url,omitempty"`
	CoverUrl      string          `json:"cover_url,omitempty"`
	FavoriteCount int64           `json:"favorite_count,omitempty"`
	CommentCount  int64           `json:"comment_count,omitempty"`
	IsFavorite    bool            `json:"is_favorite,omitempty"`
	Title         string          `json:"title,omitempty"`
}

type Comment struct {
	Id         int64           `json:"id,omitempty"`
	User       repository.User `json:"user"`
	Content    string          `json:"content,omitempty"`
	CreateDate string          `json:"create_date,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

func ConvertVideoDBToJSON(dbVideos *[]repository.Video) ([]Video, error) {
	var jsonVideos []Video
	for _, dbVideo := range *dbVideos {
		user, err := repository.NewUserDaoInstance().QueryUserByToken(dbVideo.Token)
		if err != nil {
			return jsonVideos, err
		}
		jsonVideo := Video{
			Id:            dbVideo.Id,
			Author:        *user,
			PlayUrl:       dbVideo.PlayUrl,
			CoverUrl:      dbVideo.CoverUrl,
			FavoriteCount: dbVideo.FavoriteCount,
			CommentCount:  dbVideo.CommentCount,
			IsFavorite:    false,
			Title:         dbVideo.Title,
		}
		jsonVideos = append(jsonVideos, jsonVideo)
	}

	return jsonVideos, nil
}
