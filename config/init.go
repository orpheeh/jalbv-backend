package config

import (
	"github.com/orpheeh/jalbv-backend/config/network"
)

func Init() {
	network.InitHttp()
}
