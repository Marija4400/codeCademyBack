package repo

import (
	configuration "codeCademy/configuration"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Respoitory struct {
	DB *gorm.DB
}

var repoPostres Respoitory

func CreateConnection() error {
	config := configuration.GlobalConfiguration

	databaseConfig := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		config.PostgreSQL.Host,
		config.PostgreSQL.Port,
		config.PostgreSQL.User,
		config.PostgreSQL.Password,
		config.PostgreSQL.DBName,
		config.PostgreSQL.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(databaseConfig))
	if err != nil {
		panic("Eror connecting to database, check env data! - " + err.Error())
	}

	repoPostres.DB = db
	fmt.Println("Connected to database!")
	return nil
}
