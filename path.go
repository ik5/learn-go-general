package main

import (
  "fmt"
  "os"
  "path"
)

func AppPath() (string, error) {
  if path.IsAbs(os.Args[0]) {
    return path.Dir(os.Args[0]) + string(os.PathSeparator), nil
  }
  wd, err := os.Getwd()
  if err != nil {
    return "", err
  }
  fname := path.Clean(path.Join(wd, os.Args[0]))
  return path.Dir(fname) + string(os.PathSeparator), nil

}

func main() {
  path, err := AppPath()
  if err != nil {
    panic(fmt.Sprintf("Error finding path: %s\n", err))
  } else {
    fmt.Println(path)
  }
}
