package main

import (
	"github.com/joho/godotenv"
	"github.com/orpheeh/jalbv-backend/config"
)

func main() {
	godotenv.Load()
	config.Init()
}
