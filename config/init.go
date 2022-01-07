package config

import (
	"github.com/orpheeh/jalbv-backend/config/database"
	"github.com/orpheeh/jalbv-backend/config/network"
)

func Init() {
	database.InitPostgres()
	network.InitHttp()
}
