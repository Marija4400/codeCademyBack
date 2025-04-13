package main

import (
	configuratio "codeCademy/configuration"
	initialization "codeCademy/init"
)

func main() {

	configuratio.LoadGlobalConfiguration()
	initialization.InitRedis()
}
