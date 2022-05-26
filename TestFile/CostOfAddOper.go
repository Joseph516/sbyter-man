package main

import (
	"douyin_service/global"
	"douyin_service/internal/model"
	"douyin_service/pkg/logger"
	setting2 "douyin_service/pkg/setting"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)




func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	global.Rd = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Addr,
		Password: global.RedisSetting.Password,
		DB:       0,
	})
	_,err = global.Rd.Ping().Result()
	if err!=nil{
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, // 600MB
		MaxAge:    10,  // 10天
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}


type Num struct {
	gorm.Model
	Count int64
	KeyWord string
}

func (num Num)TableName() string {
	return "test_table"
}

func (num Num) AddBySql(db *gorm.DB, key string, wg *sync.WaitGroup, mux *sync.Mutex)(err error)  {
	defer wg.Done()
	var n Num
	mux.Lock()
	db.Where("Key_word = ?",key).Debug().First(&n)
	fmt.Println(n)
	db.Model(&n).Update("count", n.Count+1)
	mux.Unlock()
	return nil

}

func (num Num) AddByRedis(rd *redis.Client, key string, wg *sync.WaitGroup)(err error)  {
	defer wg.Done()
	rd.IncrBy(key, 1)
	return nil
}
func AddBySql()  {
	mux := sync.Mutex{}
	k := "video"
	var n Num
	wg := sync.WaitGroup{}

	for i:=0;i<10;i++ {
		wg.Add(1)
		go n.AddBySql(global.DBEngine, k, &wg, &mux)
	}
	wg.Wait()
}

func AddByRedis(k string)  {

	var n Num
	wg := sync.WaitGroup{}

	for i:=0;i<100;i++ {
		wg.Add(1)
		go n.AddByRedis(global.Rd, k, &wg)
	}
	wg.Wait()
}

func main() {
	//global.DBEngine.AutoMigrate(&Num{})
	k := "test_favor"
	e := global.Rd.SetNX(k, 1, time.Minute)
	fmt.Println(e, ",", e.Val())
	//err := global.Rd.Set(k,1, time.Minute*5)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//t1 := time.Now()
	//AddByRedis(k)
	//t2 := time.Now()
	//during := t2.Nanosecond()-t1.Nanosecond()
	//fmt.Println("耗时：" ,float64(during)/float64(time.Second.Nanoseconds()))
	//count := global.Rd.Get(k)
	//fmt.Println(count)


}


