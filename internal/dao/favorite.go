package dao

import (
	"douyin_service/internal/model"
)

func (d *Dao) FavorAction(userId int64, videoId int64) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	err := favorite.Create(d.engine)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) CancelFavorAction(userId int64, videoId int64) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	err := favorite.Delete(d.engine)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) IsFavor(userId int64, videoId int64) (bool, error) {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	ok, err := favorite.IsFavor(d.engine)
	if err != nil {
		return false, err
	}
	return ok, err

}

func (d *Dao) QueryFavoritedCnt(videoId int64) (int64, error) {
	favorite := model.Favorite{
		VideoId: videoId,
	}
	cnt, err := favorite.QueryFavoritedCnt(d.engine)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (d *Dao) QueryFavoriteByUserId(userId int64) ([]int64, error) {
	favorite := model.Favorite{
		UserId: userId,
	}

	userIds, err := favorite.QueryFavoriteByUserId(d.engine)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
