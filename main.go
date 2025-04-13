package main

import (
	configuratio "codeCademy/configuration"
	initialization "codeCademy/init"
	"codeCademy/migrate"
	"codeCademy/repo"
)

func main() {

	configuratio.LoadGlobalConfiguration()
	initialization.InitRedis()
  repo.CreateConnection()
  migrate.Makemigrations()

}
