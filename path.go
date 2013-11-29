package main

import (
  "fmt"
  "os"
  "path"
)

func AppPath() string {
  if path.IsAbs(os.Args[0]) {
    return os.Args[0]
  } else {
    wd, err := os.Getwd()
    if err != nil {
         panic(fmt.Sprintf("Getwd failed: %s", err))
    }
    fname := path.Clean(path.Join(wd, os.Args[0]))
    return fname

  }
}

func main() {
  fmt.Println(AppPath())
}
