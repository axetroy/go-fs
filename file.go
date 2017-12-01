package fs

import (
  "os"
  "io/ioutil"
)

/**
ensure the file exist
 */
func EnsureFile(filepath string) (err error) {
  var (
    file *os.File
  )
  if _, err = os.Stat(filepath); os.IsNotExist(err) {
    file, err = os.Create(filepath)
    defer func() {
      file.Close()
    }()
  }
  return
}

/*
write a file
 */
func WriteFile(filename string, data []byte) error {
  return ioutil.WriteFile(filename, data, os.ModePerm)
}

/**
read a file
 */
func ReadFile(filename string) ([]byte, error) {
  return ioutil.ReadFile(filename)
}
