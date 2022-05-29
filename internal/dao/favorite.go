package dao

import (
	"douyin_service/internal/model"
)

// dao的favorite相关操作（全部走数据库）

func (d *Dao) FavorAction(userId uint, videoId uint) error {
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

func (d *Dao) CancelFavorAction(userId uint, videoId uint) error {
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

func (d *Dao) IsFavor(userId uint, videoId uint) (bool, error) {
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

func (d *Dao) QueryFavoritedCnt(videoId uint) (int64, error) {
	favorite := model.Favorite{
		VideoId: videoId,
	}
	cnt, err := favorite.QueryFavoritedCnt(d.engine)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (d *Dao) QueryFavoriteByUserId(userId uint) ([]uint, error) {
	favorite := model.Favorite{
		UserId: userId,
	}

	userIds, err := favorite.QueryFavoriteByUserId(d.engine)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
