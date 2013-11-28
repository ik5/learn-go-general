package main

import (
  "log"
  "os"
)

func main() {
 file, err := os.OpenFile("/tmp/logger.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
 if (err != nil) {
   log.Fatal("Could not create logger.log")
 }
 defer file.Close()

 log.SetOutput(file)
 log.Println("This is a test log entry")
}
