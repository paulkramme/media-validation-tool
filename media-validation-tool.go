// Copyright (C) 2017 Paul Kramme
package main

import (
  "fmt";
  "os";
  "path/filepath";
  "io";
  "crypto/md5";
  "bufio";
  "strings";
  "encoding/hex";

)

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

func scanfiles(location string) (m map[string]string, err error) {
  m = make(map[string]string)
  var walkcallback = func(path string, fileinfo os.FileInfo, inputerror error) (err error) {
    checksum,_ := md5sum(path)
    m[path] = checksum
    return
  }
  err = filepath.Walk(location, walkcallback)
  return
}

func main() {
  fmt.Println("MEDIA VALIDATION TOOL")
  fmt.Println("Copyright (C) 2017 Paul Kramme")

  if len(os.Args) > 1 {
    if os.Args[1] == "create" {
      fmt.Println("\n:: Creating media record for current directory\n")
      table, err := scanfiles(".")
      if err != nil {
        fmt.Println("Error during scan...")
      }
      //open file
      f, fileopenerror := os.Create("./media_record.csv")
      if fileopenerror != nil {
        panic(fileopenerror)
      }
      for path, hash := range table {
        if table[path] != "" {
          fmt.Fprintf(f, "%s,%s\n", path, hash)
        }
      }
    } else {
      fmt.Println("Invalid argument:", os.Args[1])
    }
  } else {
    fmt.Println("Checking file integrity")
    table, err := scanfiles(".")
    if err != nil {
      fmt.Println("Error during scan...")
    }
    fmt.Println("Checking against media_record.csv")
    file, _ := os.Open("./media_record.csv")
    scanner := bufio.NewScanner(file) //read lines
    old_data := make(map[string]string)
    for scanner.Scan() {
         splitted_string := strings.Split(scanner.Text(), ",")
         if splitted_string[1] != "" {
              old_data[splitted_string[0]] = splitted_string[1]
         }
    }
    if err != nil {
         panic(err)
    }
    var errorcount int = 0
    for path, hash := range old_data {
         if val, ok := table[path]; ok {
              if val == hash {
                   fmt.Printf("SUCCESS %s\n", path)
              } else {
                   fmt.Printf("FAIL %s\n", path)
                   errorcount++
              }
         }
    }
    fmt.Println(errorcount, "Error(s)")
    buf := bufio.NewReader(os.Stdin)
    buf.ReadBytes('\n')
  }
}
