package main

import (
  "fmt"
  "os/exec"
  "log"
)

func main() {
  cmd := exec.Command("/bin/find", "/")
  err := cmd.Start()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("All done.")
}
