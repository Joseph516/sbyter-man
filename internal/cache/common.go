package cache

import "time"

// IsExist 是否存在名为key的缓存
func (r *Redis) IsExist(key string) (bool, error) {
	result, err := r.redis.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// Set 设置k-v
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.redis.Set(key, value, expiration).Err()
}

// Get 获取key对于的value
func (r *Redis) Get(key string) (string, error) {
	return r.redis.Get(key).Result()
}

// Expire 设置（更新）key的过期时间
func (r *Redis) Expire(key string, expiration time.Duration) error {
	return r.redis.Expire(key, expiration).Err()
}

// DeleteOneCache 清除一条指定key的缓存
func (r *Redis) DeleteOneCache(key string) (bool, error) {
	err := r.redis.Del(key).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetCacheKeys 获取给定形式的key
func (r *Redis) GetCacheKeys(pattern string) ([]string, error) {
	result, err := r.redis.Keys(pattern).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
