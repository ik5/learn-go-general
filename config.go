package main

import (
  "code.google.com/p/gcfg"
  "fmt"
  "log"
)

type Config struct {
  Server struct {
    Port uint16
  }

  Exec struct {
    Path   string
    App    string
    Number string
    Queue  string
  }
}

func (conf Config) Print() {
  fmt.Println("Server:")
  fmt.Printf("\tPort: %d\n", conf.Server.Port)
  fmt.Println("Exec:")
  fmt.Printf("\tPath: %s\n",conf.Exec.Path)
  fmt.Printf("\tApp: %s\n", conf.Exec.App)
  fmt.Printf("\tNumber: %s\n", conf.Exec.Number)
  fmt.Printf("\tQueue: %s\n", conf.Exec.Queue)
}

func main() {
  var cfg Config
  err := gcfg.ReadFileInto(&cfg,"/tmp/config.ini")
  if err != nil {
    log.Fatalf("Unable to read config.init: %s", err)
  }

  cfg.Print()
}
