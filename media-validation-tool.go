


package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func sha256sum(filePath string) (result string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := sha256.New()
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
		checksum, _ := sha256sum(path)
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
				panic(err)
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
		file, err := os.Open("./media_record.csv")
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file) //read lines
		old_data := make(map[string]string)
		for scanner.Scan() {
			splitted_string := strings.Split(scanner.Text(), ",")
			if splitted_string[1] != "" {
				old_data[splitted_string[0]] = splitted_string[1]
			}
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
