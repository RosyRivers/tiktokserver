package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	Id            int64     `gorm:"column:id"`
	Token         string    `gorm:"column:token"`
	PlayUrl       string    `gorm:"column:playurl"`
	CoverUrl      string    `gorm:"column:coverurl"`
	FavoriteCount int64     `gorm:"column:favouritecount"`
	CommentCount  int64     `gorm:"column:commentcount"`
	Title         string    `gorm:"column:title"`
	LastUpTime    time.Time `gorm:"column:last_up_time"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}
func (*VideoDao) QueryVideoByAuthor(token string) (*[]Video, error) {
	var videoList []Video
	err := db.Where("token = ?", token).Find(&videoList).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, err
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &videoList, err
}
func (*VideoDao) QueryVideoLatest() (id int64, err error) {
	err = db.Table("video").Select("MAX(id)").Find(&id).Error
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return id, nil
}

func (*VideoDao) CreateVideo(video *Video) error {
	if err := db.Create(video).Error; err != nil {
		return err
	}
	return nil
}
