package main

import (
	"github.com/vgtom/vtoken/bootstrap"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
