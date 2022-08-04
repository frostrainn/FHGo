package global

import "gorm.io/gorm"

//var (
//	FHG_DB                  *gorm.DB
//	FHG_DBList              map[string]*gorm.DB
//	FHG_REDIS               *redis.Client
//	FHG_CONFIG              config.Server
//	FHG_VP                  *viper.Viper
//	FHG_LOG                 *zap.Logger
//	FHG_Timer               timer.Timer = timer.NewTimerTask()
//	FHG_Concurrency_Control             = &singleflight.Group{}
//
//	BlackCache local_cache.Cache
//	lock       sync.RWMutex
//)

var (
	DB *gorm.DB
)
