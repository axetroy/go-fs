package fs

import (
  "os"
  "io/ioutil"
  "bufio"
  "io"
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

func AppendFile(file string, data []byte) (error) {
  if f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
    return err
  } else {
    defer func() {
      err = f.Close()
    }()
    if _, err := f.Write(data); err != nil {
      return err
    }
    return nil
  }
}

func Truncate(path string, len int64) (error) {
  return os.Truncate(path, len)
}

/**
Create a read stream
 */
func CreateReadStream(path string) (stream io.Reader, err error) {
  var (
    file *os.File
  )
  if file, err = os.Open(path); err != nil {
    return
  }

  defer func() {
    err = file.Close()
  }()

  stream = bufio.NewReader(file)

  return
}

/**
Create a read stream
 */
func CreateWriteStream(path string) (stream io.Writer, err error) {
  var (
    file *os.File
  )
  if file, err = os.Open(path); err != nil {
    return
  }

  defer func() {
    err = file.Close()
  }()

  stream = bufio.NewWriter(file)

  return
}
