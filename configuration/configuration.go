package configuratio

import (
  "encoding/json"
  "fmt"
  "os"
)

var GlobalConfiguration *Configuration

type Configuration struct{
  Server struct{
    Host string `json:"Host"`
    Port string `json:"Port"`
    Mode string `json:"Mode"`
    Timeout struct {
      Read int `json:"Read"`
      Write int `json:"Write"`
    }`json:"Timeout"`

    Origin struct {
      Cors []string `json:"Cors"`
    } `json:"Origin"`
  }
  PostgreSQL struct{
    Host string `json:"Host"`
    Port int64 `json:"Port"`
    User string `json:"User"`
    Password string `json:"Password"`
    DBName string `json:"DBName"`
    SSLMode string `json:"SSLMode"`
  } `json:"PostgreSQL"`

  Redis struct{
    Address string `json:"Address"`
    Password string `json:"Password"`
    RedisDB string `json:"RedisDB"`
  }
}

func ReadConfiguration(f string) *Configuration {
  file, err := os.Open(f)
  if err!= nil{
    panic("Error opening configuration file, check if it exists first ["+err.Error()+"]")
  }

  defer func (){
    _=file.Close()
  }()

  config := &Configuration{}
  decoder:= json.NewDecoder(file)
  if err := decoder.Decode(&config); err!=nil{
    panic("File was opened but could not be decoded, Line 56 in ReadConfiguration(f string), file configuration/configuration.go")
  }
  if config.Server.Timeout.Write < 1{
    config.Server.Timeout.Write=15
  }

  if config.Server.Timeout.Read < 1{
    config.Server.Timeout.Read = 15
  }

  fmt.Println("Configuration is set!")
  fmt.Println(config)

  return config
}

func LoadGlobalConfiguration()  {

  config := ReadConfiguration(os.Args[1])
  GlobalConfiguration = config
  
}














