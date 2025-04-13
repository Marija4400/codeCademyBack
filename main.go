package main

import (
	configuratio "codeCademy/configuration"
	initialization "codeCademy/init"
	"codeCademy/repo"
)

func main() {

	configuratio.LoadGlobalConfiguration()
	initialization.InitRedis()
  repo.CreateConnection()
}
