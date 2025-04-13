package repository

import "gorm.io/gorm"


type Respoitory struct{
  DB *gorm.DB
}

var repoPostres Respoitory


fucn CreateConnection() error{
  
}
