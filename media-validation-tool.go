// Copyright (C) 2017 Paul Kramme
package main

import (
  "fmt";
  "os";
  "path/filepath";
  "io";
  "crypto/md5";
  "encoding/hex"
)

//type WalkCallb func(path string, info os.FileInfo, err error) error

func md5sum(filePath string) (result string, err error) {
  file, err := os.Open(filePath)
  if err != nil {
    return
  }
  defer file.Close()

  hash := md5.New()
  _, err = io.Copy(hash, file)
  if err != nil {
    return
  }

  result = hex.EncodeToString(hash.Sum(nil))
  return
}

func scanfiles(location string) (m map[int]string, err error) {
  var walkcallback = func(path string, fileinfo os.FileInfo, inputerror error) (err error) {
    checksum,_ := md5sum(path)
    fmt.Println(path, checksum)
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
      fmt.Println("\n:: Creating media record for current directory\n")
      scanfiles(".")
    } else {
      fmt.Println("Invalid argument:", os.Args[1])
    }
  }
}
