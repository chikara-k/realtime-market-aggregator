package initializer

import (
	"github.com/chikara-k/realtime-market-aggregator/config"
	"github.com/chikara-k/realtime-market-aggregator/infra/db/inframysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func initDBConnection(conf config.Config) (*inframysql.RealtimeMarketAggregatorDBConnection, error) {
	dsn := conf.DB.GetSettingStr()

	glog := logger.Default.LogMode(logger.Warn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: glog})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return inframysql.NewRealtimeMarketAggregatorDBConnection(db), err
}

func InitDBConnection(conf config.Config) *inframysql.RealtimeMarketAggregatorDBConnection {
	db, err := initDBConnection(conf)
	if err != nil {
		panic(err)
	}
	return db
}
