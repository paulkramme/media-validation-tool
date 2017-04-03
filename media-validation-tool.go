// Copyright (C) 2017 Paul Kramme
package main

import (
  "fmt";
  "os";
  "path/filepath"
)

//type WalkCallb func(path string, info os.FileInfo, err error) error

func scanfiles(location string) (m map[int]string, err error) {
  var walkcallback = func(path string, fileinfo os.FileInfo, inputerror error) (err error) {
    fmt.Println(path)
    return
  }
  var something map[int]string
  err = filepath.Walk(location, walkcallback)
  return something, err
}

func main() {
  fmt.Println("MEDIA VALIDATION TOOL")
  fmt.Println("Copyright (C) 2017 Paul Kramme")

  if len(os.Args) > 1 {
    if os.Args[1] == "create" {
      fmt.Println("\n:: Creating media record for current directory")
      scanfiles(".")
    } else {
      fmt.Println("Invalid argument:", os.Args[1])
    }
  }
}
