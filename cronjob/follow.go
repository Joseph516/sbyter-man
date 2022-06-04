package cronjob

import (
	"douyin_service/global"
	"douyin_service/internal/dao"
	"douyin_service/pkg/util"
	"strconv"
)

const (
	FOLLOWCNTTIME = "0 30 2 1/1 * *" //每日的凌晨2点半执行一次
)

func (dc DouyinCron)FlashFollowCnt()  {
	global.Logger.Info("follow count flash begin...")
	pattern := util.FOLLOWCOUNTFREFIX+"*"
	keys, err := global.Rd.Keys(pattern).Result()
	if err !=nil{
		global.Logger.Error("flash follow count error:", err)
	}
	for _ , key := range keys{
		val,err := global.Rd.Get(key).Result()
		if err!=nil{
			global.Logger.Error(err)
		}
		id, err := strconv.Atoi(key[len(util.FOLLOWCOUNTFREFIX):])
		if err!=nil{
			global.Logger.Error(err)
		}
		d := dao.New(global.DBEngine)

		user, err := d.GetUserById(uint(id))
		if err!=nil{
			global.Logger.Error(err)
		}
		cnt, _ := strconv.Atoi(val)
		user.FollowCount = int64(cnt)
		user.Update(global.DBEngine)
	}
	global.Logger.Info("follow count flash end")
}

func (dc DouyinCron)FlashFanCnt()  {
	global.Logger.Info("fan count flash begin...")
	pattern := util.FANCOUNTFREFIX+"*"
	keys, err := global.Rd.Keys(pattern).Result()
	if err !=nil{
		global.Logger.Error("flash fan count error:", err)
	}
	for _ , key := range keys{
		val,err := global.Rd.Get(key).Result()
		if err!=nil{
			global.Logger.Error(err)
		}
		id, err := strconv.Atoi(key[len(util.FANCOUNTFREFIX):])
		if err!=nil{
			global.Logger.Error(err)
		}
		d := dao.New(global.DBEngine)

		user, err := d.GetUserById(uint(id))
		if err!=nil{
			global.Logger.Error(err)
		}
		cnt, _ := strconv.Atoi(val)
		user.FollowerCount = int64(cnt)
		user.Update(global.DBEngine)
	}
	global.Logger.Info("fan count flash end")

}