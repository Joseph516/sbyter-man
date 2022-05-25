package cache

import (
	"douyin_service/pkg/util"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// FollowAction userId给关注指定的up
// 必须保证此时缓存中有对应的数据
func (r *Redis) FollowAction(upId uint, fanId uint) (bool, string, error) {
	upKey := util.FanCountKey(upId)
	fanKey := util.FollowCountKey(fanId)
	if flag, _:=r.IsExist(upKey); !flag{
		// 缓存中没有
		return false, upKey, errors.New("缓存中没有项:" + upKey)
	}
	if flag, _ := r.IsExist(fanKey);!flag{
		// 缓存中没有
		return false, fanKey, errors.New("缓存中没有项:" + fanKey)
	}
	r.IncrFanCnt(upId)
	r.IncrFollowCnt(fanId)
	return true, "", nil
}

// CancelFollowAction userId取消关注指定up
// 必须保证缓存中有该条数据
func (r *Redis) CancelFollowAction(upId uint, fanId uint) (bool, string, error) {
	upKey := util.FanCountKey(upId)
	fanKey := util.FollowCountKey(fanId)
	if flag, _:=r.IsExist(upKey); !flag{
		// 缓存中没有
		return false, upKey, errors.New("缓存中没有项:" + upKey)
	}
	if flag, _ := r.IsExist(fanKey);!flag{
		// 缓存中没有
		return false, fanKey, errors.New("缓存中没有项:" + fanKey)
	}
	r.DecrFanCnt(upId)
	r.DecrFollowCnt(fanId)
	return true, "", nil
}

// QueryFollowCnt 查询某用户关注的up数量
func (r *Redis) QueryFollowCnt(userId uint) (bool, int64, error) {
	key := util.FollowCountKey(userId)
	exist, err := r.IsExist(key)
	if err != nil {
		return false, 0, err
	}
	if !exist {
		return false, 0, nil
	}
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return false, 0, err
	}
	cnt, _ := strconv.Atoi(result)
	return true, int64(cnt), nil
}

// IncrFollowCnt 增加关注up数量
func (r *Redis) IncrFollowCnt(userId uint) int64 {
	key := util.FollowCountKey(userId)
	result := r.redis.Incr(key).Val()
	return result
}

// DecrFollowCnt 减少关注的up数量
func (r *Redis) DecrFollowCnt(userId uint) int64 {
	key := util.FollowCountKey(userId)
	result := r.redis.Decr(key).Val()
	return result
}

// QueryFanCnt 查询某up的粉丝数量
func (r *Redis) QueryFanCnt(userId uint) (bool, int64, error) {
	key := util.FanCountKey(userId)
	exist, err := r.IsExist(key)
	if err != nil {
		return false, 0, err
	}
	if !exist {
		return false, 0, nil
	}
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return false, 0, err
	}
	cnt, _ := strconv.Atoi(result)
	return true, int64(cnt), nil
}

// IncrFanCnt  增加up的粉丝数量并刷新生存周期
func (r *Redis) IncrFanCnt(userId uint) int64 {
	key := util.FanCountKey(userId)
	result := r.redis.Incr(key).Val()
	r.redis.Expire(key, time.Hour*24)
	return result
}

// DecrFanCnt  减少up的粉丝数量并刷新生存周期
func (r *Redis) DecrFanCnt(userId uint) int64 {
	key := util.FanCountKey(userId)
	result := r.redis.Decr(key).Val()
	r.redis.Expire(key, time.Hour*24)
	return result
}

func (r *Redis) SetFollowInfo(key string, val int64){
	r.Set(key, val, time.Hour*24)
}


func (r *Redis) Output(){
	vals, _ :=r.redis.Keys("*_COUNT").Result()
	for i:=range vals{
		k := vals[i]
		v := r.redis.Get(k)
		fmt.Println(k,":",v)
	}
}