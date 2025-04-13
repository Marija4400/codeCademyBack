package migrate

import(
  "codeCademy/models"
  "codeCademy/repo"
)

func Makemigrations(){
  repo.RepoPostres.DB.AutoMigrate(
    &models.User{},
    &models.Role{},
  )
}
