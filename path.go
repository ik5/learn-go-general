package main

import (
  "fmt"
  "os"
  "path"
  "runtime"
)

func AppPath() (string, error) {
  separator := func() string {
    switch runtime.GOOS {
      case "windows" :
        return "\\"
      case "linux" :
        return "/"

      default :
        return "/"
    }

  }
  if path.IsAbs(os.Args[0]) {
    return path.Dir(os.Args[0]) + separator(), nil
  }
  wd, err := os.Getwd()
  if err != nil {
    return "", err
  }
  fname := path.Clean(path.Join(wd, os.Args[0]))
  return path.Dir(fname) + separator(), nil

}

func main() {
  path, err := AppPath()
  if err != nil {
    panic(fmt.Sprintf("Error finding path: %s\n", err))
  } else {
    fmt.Println(path)
  }
}
