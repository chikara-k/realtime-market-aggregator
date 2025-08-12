package main

import (
	"github.com/chikara-k/realtime-market-aggregator/config"
	"github.com/chikara-k/realtime-market-aggregator/initializer"
)

func main() {
	conf := config.LoadConfig()

	db := initializer.InitDBConnection(conf)
	defer func() {
		sqlDB, err := db.Conn.DB()
		if  err != nil {
			panic(err)
		}
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}()
}