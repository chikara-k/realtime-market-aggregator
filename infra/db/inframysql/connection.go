package inframysql

import "gorm.io/gorm"

type RealtimeMarketAggregatorDBConnection struct {
	Conn *gorm.DB
}

func NewRealtimeMarketAggregatorDBConnection(conn *gorm.DB) *RealtimeMarketAggregatorDBConnection {
	return &RealtimeMarketAggregatorDBConnection{Conn: conn}
}
